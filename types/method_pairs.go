package types

import "strings"

type Pair struct {
	Token0Address string `json:"token0_address"`
	Token1Address string `json:"token1_address"`
	Fee           int64  `json:"fee"`
	TickSpacing   int64  `json:"tick_spacing"`
	PoolAddress   string `json:"pool_address"`
	PoolType      uint8  `json:"pool_type"`
	CreatedAt     int64  `json:"created_at"`
	Hash          string `json:"hash"`
	ChainID       int16  `json:"chain_id"`
}

func (p *Pair) Lower() {
	p.Token0Address = strings.ToLower(p.Token0Address)
	p.Token1Address = strings.ToLower(p.Token1Address)
	p.PoolAddress = strings.ToLower(p.PoolAddress)
	p.Hash = strings.ToLower(p.Hash)
}
