package types

type StateChangeView struct {
	Cause map[string]interface{} `json:"cause"`
	Value map[string]interface{} `json:"value"`
}
