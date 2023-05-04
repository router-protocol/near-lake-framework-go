package types

type ExecutionOutcomeView struct {
	Logs        []string              `json:"logs"`
	ReceiptIds  []string              `json:"receipt_ids"`
	GasBurnt    uint64                `json:"gas_burnt"`
	TokensBurnt *BigInt               `json:"tokens_burnt"`
	ExecutorId  string                `json:"executor_id"`
	Status      Status                `json:"status"`
	Metadata    ExecutionMetadataView `json:"metadata"`
}

type Status map[string]interface{}

func (status Status) IsUnknown() bool {
	_, ok := status["Unknown"]
	return ok
}

func (status Status) IsFailure() bool {
	_, ok := status["Failure"]
	return ok
}

func (status Status) IsSuccess() bool {
	_, ok1 := status["SuccessValue"]
	_, ok2 := status["SuccessReceiptId"]

	return ok1 || ok2
}

func (status Status) SuccessValue() *string {
	_, ok := status["SuccessValue"]
	if status.IsSuccess() && ok {
		return status["SuccessValue"].(*string)
	}
	return nil
}

func (status Status) SuccessReceiptId() *string {
	_, ok := status["SuccessReceiptId"]
	if status.IsSuccess() && ok {
		return status["SuccessReceiptId"].(*string)
	}
	return nil
}
