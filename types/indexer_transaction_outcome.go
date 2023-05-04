package types

type IndexerTransactionWithOutcome struct {
	Transaction SignedTransactionView                      `json:"transaction"`
	Outcome     IndexerExecutionOutcomeWithOptionalReceipt `json:"outcome"`
}
