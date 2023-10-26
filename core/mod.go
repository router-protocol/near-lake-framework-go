package core

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/router-protocol/near-lake-framework-go/types"
)

type LakeConfig struct {
	s3BucketName          string
	s3RegionName          string
	startBlockHeight      uint64
	s3Config              *client.ConfigProvider
	blocksPreloadPoolSize uint64
}

func Streamer(ctx context.Context, config LakeConfig, numOfWorkers int) chan types.StreamMessage {
	fmt.Println("Starting Streamer...")
	messageChannel := start(ctx, config, numOfWorkers)
	return messageChannel
}

func start(ctx context.Context, config LakeConfig, numOfWorkers int) chan types.StreamMessage {
	messageChannel := make(chan types.StreamMessage)
	// closeSignal := make(chan bool)

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String(config.s3RegionName),
	})
	s3Client := s3.New(awsSession)
	if config.s3Config != nil {
		s3Client = s3.New(*config.s3Config)
	}
	s3Fetcher := S3Fetcher{}

	go func() {
		blocks, err := s3Fetcher.ListBlocks(s3Client, config.s3BucketName, config.startBlockHeight, config.blocksPreloadPoolSize)
		if err != nil {
			fmt.Println("1", err)
			return
		}
		if len(blocks) == 0 {
			return
		}
		numWorkers := numOfWorkers
		if len(blocks) < numWorkers {
			numWorkers = len(blocks)
		}

		startTime := time.Now()
		blockHeightsPerWorker := (len(blocks) + numWorkers - 1) / numWorkers

		var wg sync.WaitGroup
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(workerIndex int) {
				defer wg.Done()
				startIndex := workerIndex * blockHeightsPerWorker
				endIndex := ((workerIndex + 1) * blockHeightsPerWorker) - 1

				if startIndex > len(blocks)-1 {
					return
				}
				if endIndex > len(blocks)-1 {
					endIndex = len(blocks) - 1
				}
				for _, blockHeight := range blocks[startIndex : endIndex+1] {
					message, err := s3Fetcher.FetchStreamerMessage(s3Client, config.s3BucketName, blockHeight)
					if err != nil {
						fmt.Println(err)
						continue
					}

					select {
					case messageChannel <- *message:
					case <-ctx.Done():
						return
					}
				}
			}(i)
		}
		wg.Wait()
		close(messageChannel)

		endTime := time.Now()
		fmt.Printf("Processed %d blocks in %v\n", len(blocks), endTime.Sub(startTime))
		fmt.Println("Streamer ended.")
	}()

	return messageChannel
}
