package core

import (
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

func Streamer(config LakeConfig, numWorkers int) chan types.StreamMessage {
	fmt.Println("Starting Streamer...")
	messageChannel := make(chan types.StreamMessage, numWorkers)
	go func(cfg LakeConfig, mc chan types.StreamMessage, workers int) {
		start(cfg, mc, workers)
		fmt.Println("Streamer ended.")
	}(config, messageChannel, numWorkers)
	return messageChannel
}

func start(config LakeConfig, messageChannel chan types.StreamMessage, numWorkers int) {
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	s3Client := s3.New(awsSession)
	if config.s3Config != nil {
		s3Client = s3.New(*config.s3Config)
	}
	s3Fetcher := S3Fetcher{}

	blocks, err := s3Fetcher.ListBlocks(s3Client, config.s3BucketName, config.startBlockHeight, config.blocksPreloadPoolSize)
	if err != nil {
		fmt.Println("1", err)
		return
	}
	//if workers are greater than number of blocks
	if len(blocks) <= numWorkers {
		numWorkers = len(blocks)
	}
	// fmt.Println("len of Blocks:", len(blocks), numWorkers)
	startTime := time.Now()
	blockHeightsPerWorker := (len(blocks) + numWorkers - 1) / numWorkers
	// fmt.Println("blockHeightsPerWorker", blockHeightsPerWorker)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerIndex int) {
			defer wg.Done()
			startIndex := workerIndex * blockHeightsPerWorker
			endIndex := ((workerIndex + 1) * blockHeightsPerWorker) - 1
			//When no more blocks remain to assign go routine
			if startIndex > len(blocks)-1 {
				// fmt.Println("no block left")
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
				messageChannel <- *message
			}
		}(i)
	}
	wg.Wait()
	close(messageChannel)

	endTime := time.Now()
	fmt.Printf("Processed %d blocks in %v\n", len(blocks), endTime.Sub(startTime))
}
