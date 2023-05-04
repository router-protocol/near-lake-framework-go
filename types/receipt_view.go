package types

import "encoding/json"

type ReceiptView struct {
	PredecessorId string  `json:"predecessor_id"`
	ReceiverId    string  `json:"receiver_id"`
	ReceiptId     string  `json:"receipt_id"`
	Receipt       Receipt `json:"receipt"`
}

type DataReceiverView struct {
	DataId     string `json:"data_id"`
	ReceiverId string `json:"receiver_id"`
}

type Data struct {
	DataId string
	Data   []uint8
}

type Action struct {
	SignerId            string             `json:"signer_id"`
	SignerPublicKey     string             `json:"signer_public_key"`
	GasPrice            *BigInt            `json:"gas_price"`
	OutputDataReceivers []DataReceiverView `json:"output_data_receivers"`
	InputDataIds        []string           `json:"input_data_ids"`
	Actions             []ActionView       `json:"actions"`
}

type Receipt map[string]interface{}

func (receipt Receipt) IsAction() bool {
	_, ok := receipt["Action"]
	return ok
}

func (receipt Receipt) IsData() bool {
	_, ok := receipt["Data"]
	return ok
}

func (receipt Receipt) GetAction() *Action {
	if receipt.IsAction() {
		data, err := json.Marshal(receipt["Action"])
		if err != nil {
			return nil
		}
		action := Action{}
		err = json.Unmarshal(data, &action)
		return &action
	}
	return nil
}

func (receipt Receipt) GetData() *Data {
	if receipt.IsData() {
		return receipt["Data"].(*Data)
	}
	return nil
}
