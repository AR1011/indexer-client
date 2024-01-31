package types

type GetBlockTimestampsParams struct {
	ChainID   int64
	FromBlock int64
	ToBlock   int64
}

type GetBlockAtTimestampParams struct {
	ChainID   int64
	Timestamp int64
}

type FindTokensParams struct {
	ChainID int64
	Filter  *TokenFilter
	Options *TokenOptions
}

type FindPairsParams struct {
	ChainID int64
	Filter  *PairFilter
	Options *PairOptions
}
