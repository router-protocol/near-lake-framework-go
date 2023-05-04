package types

type StreamMessage struct {
	Block  BlockView      `json:"block"`
	Shards []IndexerShard `json:"shards"`
}
