package types

type BlockTimestamp struct {
	Block     int64 `json:"block"`
	Timestamp int64 `json:"timestamp"`
}
type Heights struct {
	Blocks int64
	Tokens int64
	Pairs  int64
}
