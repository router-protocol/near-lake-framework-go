package types

type ExecutionMetadataView struct {
	Version    uint32        `json:"version"`
	GasProfile []CostGasUsed `json:"gas_profile"`
}

type CostGasUsed struct {
	CostCategory string  `json:"cost_category"`
	Cost         string  `json:"cost"`
	GasUsed      *BigInt `json:"gas_used"`
}
