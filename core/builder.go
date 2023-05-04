package core

import (
	"github.com/aws/aws-sdk-go/aws/client"
)

type LakeConfigBuilder struct {
	lakeConfig *LakeConfig
}

type ILakeConfigBuilder interface {
	SetS3BucketName(name string) ILakeConfigBuilder
	SetS3RegionName(name string) ILakeConfigBuilder
	SetStartBlockHeight(height uint64) ILakeConfigBuilder
	SetS3Config(config *client.ConfigProvider) ILakeConfigBuilder
	SetBlocksPreloadPoolSize(size uint64) ILakeConfigBuilder
	Mainnet() ILakeConfigBuilder
	Testnet() ILakeConfigBuilder
	Build() *LakeConfig
}

func DefaultLakeConfigBuilder() ILakeConfigBuilder {
	return &LakeConfigBuilder{
		lakeConfig: &LakeConfig{},
	}
}

const (
	MainnetBucketName = "near-lake-data-mainnet"
	TestnetBucketName = "near-lake-data-testnet"
	AwsRegion         = "eu-central-1"
)

func (cf *LakeConfigBuilder) SetS3BucketName(name string) ILakeConfigBuilder {
	cf.lakeConfig.s3BucketName = name
	return cf
}

func (cf *LakeConfigBuilder) SetS3RegionName(name string) ILakeConfigBuilder {
	cf.lakeConfig.s3RegionName = name
	return cf
}

func (cf *LakeConfigBuilder) SetStartBlockHeight(height uint64) ILakeConfigBuilder {
	cf.lakeConfig.startBlockHeight = height
	return cf
}

func (cf *LakeConfigBuilder) SetS3Config(config *client.ConfigProvider) ILakeConfigBuilder {
	cf.lakeConfig.s3Config = config
	return cf
}

func (cf *LakeConfigBuilder) SetBlocksPreloadPoolSize(size uint64) ILakeConfigBuilder {
	cf.lakeConfig.blocksPreloadPoolSize = size
	return cf
}

func (cf *LakeConfigBuilder) Mainnet() ILakeConfigBuilder {
	cf.lakeConfig.s3RegionName = AwsRegion
	cf.lakeConfig.s3BucketName = MainnetBucketName
	return cf
}

func (cf *LakeConfigBuilder) Testnet() ILakeConfigBuilder {
	cf.lakeConfig.s3RegionName = AwsRegion
	cf.lakeConfig.s3BucketName = TestnetBucketName
	return cf
}

func (cf *LakeConfigBuilder) Build() *LakeConfig {
	return cf.lakeConfig
}
