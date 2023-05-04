package types

type IndexerChunkView struct {
	Author       string                          `json:"author"`
	Header       ChunkHeaderView                 `json:"header"`
	Transactions []IndexerTransactionWithOutcome `json:"transactions"`
	Receipts     []ReceiptView                   `json:"receipts"`
}
