package types

type IndexerExecutionOutcomeWithOptionalReceipt struct {
	ExecutionOutcome ExecutionOutcomeWithIdView `json:"execution_outcome"`
	Receipt          ReceiptView                `json:"receipt"`
}
