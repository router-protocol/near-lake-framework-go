package types

type ChunkHeaderView struct {
	ChunkHash            string  `json:"chunk_hash"`
	PrevBlockHash        string  `json:"prev_block_hash"`
	OutcomeRoot          string  `json:"outcome_root"`
	PrevStateRoot        string  `json:"prev_state_root"`
	EncodedMerkleRoot    string  `json:"encoded_merkle_root"`
	EncodedLength        uint64  `json:"encoded_length"`
	HeightCreated        uint64  `json:"height_created"`
	HeightIncluded       uint64  `json:"height_included"`
	ShardId              uint64  `json:"shard_id"`
	GasUsed              uint64  `json:"gas_used"`
	GasLimit             uint64  `json:"gas_limit"`
	RentPaid             *BigInt `json:"rent_paid"`
	ValidatorReward      *BigInt `json:"validator_reward"`
	BalanceBurnt         *BigInt `json:"balance_burnt"`
	OutgoingReceiptsRoot string  `json:"outgoing_receipts_root"`
	TxRoot               string  `json:"tx_root"`
	Signature            string  `json:"signature"`
}
