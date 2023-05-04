package types

type BlockView struct {
	Author string            `json:"author"`
	Header BlockHeaderView   `json:"header"`
	Chunks []ChunkHeaderView `json:"chunks"`
}
