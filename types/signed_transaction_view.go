package types

type SignedTransactionView struct {
	SignerId   string        `json:"signer_id"`
	PublicKey  string        `json:"public_key"`
	Nonce      uint64        `json:"nonce"`
	ReceiverId string        `json:"receiver_id"`
	Actions    []interface{} `json:"actions"`
	Signature  string        `json:"signature"`
	Hash       string        `json:"hash"`
}
