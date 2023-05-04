package types

type IndexerShard struct {
	ShardId                  uint64                               `json:"shard_id"`
	Chunk                    *IndexerChunkView                    `json:"chunk"`
	ReceiptExecutionOutcomes []IndexerExecutionOutcomeWithReceipt `json:"receipt_execution_outcomes"`
	StateChanges             []StateChangeView                    `json:"state_changes"`
}
