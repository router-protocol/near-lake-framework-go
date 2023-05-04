package types

type ReceiptEnumView struct {
	SignerId            string             `json:"signer_id"`
	SignerPublicKey     string             `json:"signer_public_key"`
	GasPrice            *BigInt            `json:"gas_price"`
	OutputDataReceivers []DataReceiverView `json:"output_data_receivers"`
	InputDataIds        []string           `json:"input_data_ids"`
	Actions             []ActionView       `json:"actions"`
	DataId              string             `json:"data_id"`
	Data                []byte             `json:"data"`
}
