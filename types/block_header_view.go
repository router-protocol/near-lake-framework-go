package types

type BlockHeaderView struct {
	Height            uint64  `json:"height"`
	PrevHeight        *uint64 `json:"prev_height"`
	EpochId           string  `json:"epoch_id"`
	NextEpochId       string  `json:"next_epoch_id"`
	Hash              string  `json:"hash"`
	PrevHash          string  `json:"prev_hash"`
	PrevStateRoot     string  `json:"prev_state_root"`
	ChunkReceiptsRoot string  `json:"chunk_receipts_root"`
	ChunkHeadersRoot  string  `json:"chunk_headers_root"`
	ChunkTxRoot       string  `json:"chunk_tx_root"`
	OutcomeRoot       string  `json:"outcome_root"`
	ChunksIncluded    uint64  `json:"chunks_included"`
	ChallengesRoot    string  `json:"challenges_root"`
	Timestamp         uint64  `json:"timestamp"`
	TimestampNanosec  uint64  `json:"timestamp_nanosec,string"`
	RandomValue       string  `json:"random_value"`
	ChunkMask         []bool  `json:"chunk_mask"`
	GasPrice          *BigInt `json:"gas_price"`
	BlockOrdinal      *uint64 `json:"block_ordinal"`
	RentPaid          *BigInt `json:"rent_paid"`
	ValidatorReward   *BigInt `json:"validator_reward"`
	TotalSupply       *BigInt `json:"total_supply"`
	LastFinalBlock    string  `json:"last_final_block"`
	LastDsFinalBlock  string  `json:"last_ds_final_block"`
	NextBpHash        string  `json:"next_bp_hash"`
	BlockMerkleRoot   string  `json:"block_merkle_root"`
}
