# Near lake framework on Golang
## How to install
```
go get github.com/router-protocol/near-lake-framework-go
```

## How to use
```golang
config := core.DefaultLakeConfigBuilder().
		      Mainnet().
	          SetStartBlockHeight(79075963).
	          SetBlocksPreloadPoolSize(100).
	          Build()
channel := core.Streamer(*config)
for {
    select {
    case message := <-channel:
        fmt.Println(message.Shards[0].Chunk.Receipts)
    }
}
```