package core

import (
	"fmt"
	"sync"

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
	messageChannel := make(chan types.StreamMessage, 1)
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
		fmt.Println(err)
		return
	}
	fmt.Println("len of Blocks:", len(blocks), config.blocksPreloadPoolSize)

	blockHeightsPrefixes := make(chan uint64, len(blocks))
	for _, blockHeight := range blocks {
		blockHeightsPrefixes <- blockHeight
	}
	close(blockHeightsPrefixes)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for blockHeight := range blockHeightsPrefixes {
				message, err := s3Fetcher.FetchStreamerMessage(s3Client, config.s3BucketName, blockHeight)
				if err != nil {
					fmt.Println(err)
					continue
				}
				messageChannel <- *message
			}
		}()
	}
	wg.Wait()
	close(messageChannel)
}
