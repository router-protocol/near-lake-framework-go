package types

type IndexerExecutionOutcomeWithReceipt struct {
	ExecutionOutcome ExecutionOutcomeWithIdView `json:"execution_outcome"`
	Receipt          ReceiptView                `json:"receipt"`
}
