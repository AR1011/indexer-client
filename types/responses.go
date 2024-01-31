package types

import "encoding/json"

type JRPCResponse struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *JRPCError      `json:"error,omitempty"`
}

type JRPCError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type GetBlockNumberResponse struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	Result map[int64]int64 `json:"result,omitempty"`
	Error  *JRPCError      `json:"error,omitempty"`
}

type GetHeightsResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result *Heights   `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type GetChainsResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result []Chain    `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type GetBlockTimestampsResponse struct {
	ID     string            `json:"id"`
	Method string            `json:"method"`
	Result []*BlockTimestamp `json:"result,omitempty"`
	Error  *JRPCError        `json:"error,omitempty"`
}

type GetBlockAtTimestampResponse struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	Result *BlockTimestamp `json:"result,omitempty"`
	Error  *JRPCError      `json:"error,omitempty"`
}

type FindTokensResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result []*Token   `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type FindPairsResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result []*Pair    `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type GetTokenCountResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result int64      `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type GetPairCountResponse struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Result int64      `json:"result,omitempty"`
	Error  *JRPCError `json:"error,omitempty"`
}

type Chain struct {
	ChainID       int    `json:"chain_id"`
	Name          string `json:"name"`
	ShortName     string `json:"short_name"`
	ExplorerURL   string `json:"explorer_url"`
	RouterV2      string `json:"router_v2"`
	FactoryV2     string `json:"factory_v2"`
	RouterV3      string `json:"router_v3"`
	FactoryV3     string `json:"factory_v3"`
	BlockDuration int64  `json:"block_duration"`
	Http          string `json:"-"`
}
