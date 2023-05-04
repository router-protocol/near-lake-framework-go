package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"near-event-listener/near-lake-framework-go/types"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	EstimatedShardsCount = 4
	MaxRetryCount        = 10
)

type S3Fetcher struct{}

type IFetcher interface {
	ListBlocks(
		s3Client *s3.S3,
		s3BucketName string,
		startFromBlockHeight uint64,
		numberOfBlocksRequested uint64,
	) (map[string]interface{}, error)
}

func (s3fetcher *S3Fetcher) ListBlocks(
	s3Client *s3.S3,
	s3BucketName string,
	startFromBlockHeight uint64,
	numberOfBlocksRequested uint64,
) ([]uint64, error) {
	fmt.Println("aws.Int64(int64(numberOfBlocksRequested * (1 + EstimatedShardsCount))): ", aws.Int64(int64(numberOfBlocksRequested*(1+EstimatedShardsCount))))
	response, err := s3Client.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket:       aws.String(s3BucketName),
			MaxKeys:      aws.Int64(int64(numberOfBlocksRequested * (EstimatedShardsCount))),
			Delimiter:    aws.String("/"),
			StartAfter:   aws.String(fmt.Sprintf("%012d", startFromBlockHeight)),
			RequestPayer: aws.String(s3.RequestPayerRequester),
		},
	)
	if err != nil {
		return nil, err
	}
	var ret []uint64
	for _, prefixString := range response.CommonPrefixes {
		prefixes := strings.Split(*prefixString.Prefix, "/")
		for _, prefix := range prefixes {
			value, err := strconv.ParseUint(prefix, 10, 32)
			if err == nil {
				ret = append(ret, value)
			}
		}
	}
	return ret, nil
}

func (s3fetcher *S3Fetcher) FetchStreamerMessage(
	s3Client *s3.S3,
	s3BucketName string,
	blockHeight uint64,
) (*types.StreamMessage, error) {
	var message = types.StreamMessage{}
	blockViewResponse, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket:       aws.String(s3BucketName),
		Key:          aws.String(fmt.Sprintf("%012d/block.json", blockHeight)),
		RequestPayer: aws.String(s3.RequestPayerRequester),
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var bodyBytes = new(bytes.Buffer)
	_, err = bodyBytes.ReadFrom(blockViewResponse.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var block = types.BlockView{}
	err = json.Unmarshal(bodyBytes.Bytes(), &block)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	message.Block = block
	wg := &sync.WaitGroup{}
	for _, chunk := range block.Chunks {
		// Todo: Sort the result in batch request case
		//wg.Add(1)
		//go func(shardId uint64) {
		//	defer wg.Done()
		shard, err := s3fetcher.FetchShardOrRetry(s3Client, s3BucketName, blockHeight, chunk.ShardId)
		if err != nil {
			fmt.Println(err)
		} else {
			message.Shards = append(message.Shards, *shard)
		}
		//}(chunk.ShardId)
	}
	wg.Wait()
	return &message, nil
}

func (s3fetcher *S3Fetcher) FetchShardOrRetry(
	s3Client *s3.S3,
	s3BucketName string,
	blockHeight uint64,
	shardId uint64,
) (*types.IndexerShard, error) {
	var totalAttempts = 0
	var shard = types.IndexerShard{}
	for totalAttempts < MaxRetryCount {
		shardViewResponse, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket:       aws.String(s3BucketName),
			Key:          aws.String(fmt.Sprintf("%012d/shard_%d.json", blockHeight, shardId)),
			RequestPayer: aws.String(s3.RequestPayerRequester),
		})
		if err != nil {
			fmt.Println(err)
			totalAttempts++
			continue
		}
		var bodyBytes = new(bytes.Buffer)
		_, err = bodyBytes.ReadFrom(shardViewResponse.Body)
		if err != nil {
			totalAttempts++
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		err = json.Unmarshal(bodyBytes.Bytes(), &shard)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &shard, nil
	}
	return nil, nil
}
