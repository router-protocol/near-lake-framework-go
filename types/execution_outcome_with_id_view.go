package types

type ExecutionOutcomeWithIdView struct {
	Proof     []interface{}        `json:"proof"`
	BlockHash string               `json:"block_hash"`
	Id        string               `json:"id"`
	Outcome   ExecutionOutcomeView `json:"outcome"`
}
