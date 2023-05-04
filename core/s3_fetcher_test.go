package core

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestFetchStreamerMessage(t *testing.T) {
	config := DefaultLakeConfigBuilder().
		Testnet().
		SetStartBlockHeight(122033824).
		SetBlocksPreloadPoolSize(100).
		Build()
	// Create an AWS session

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	s3Client := s3.New(awsSession)
	if config.s3Config != nil {
		s3Client = s3.New(*config.s3Config)
	}
	for {
		s3Fetcher := S3Fetcher{}
		blockHeightsPrefixes, err := s3Fetcher.ListBlocks(s3Client, "near-lake-data-testnet", 122033824, 100)
		if err != nil {
			fmt.Println(err)
		}
		if len(blockHeightsPrefixes) <= 0 {
			continue
		}

		for _, blockHeight := range blockHeightsPrefixes {
			message, err := s3Fetcher.FetchStreamerMessage(s3Client, config.s3BucketName, blockHeight)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(message.Block.Header.Height)
		}
	}

}

func TestListBlocks(t *testing.T) {
	config := DefaultLakeConfigBuilder().
		Testnet().
		SetStartBlockHeight(122033824).
		SetBlocksPreloadPoolSize(100).
		Build()
	// Create an AWS session

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	s3Client := s3.New(awsSession)
	if config.s3Config != nil {
		s3Client = s3.New(*config.s3Config)
	}
	for {
		s3Fetcher := S3Fetcher{}
		blockHeightsPrefixes, err := s3Fetcher.ListBlocks(s3Client, "near-lake-data-testnet", 122033824, 100)
		if err != nil {
			fmt.Println(err)
		}
		if len(blockHeightsPrefixes) <= 0 {
			continue
		}
		fmt.Println("length of blockHeightsPrefixes:", len(blockHeightsPrefixes))
	}

}
