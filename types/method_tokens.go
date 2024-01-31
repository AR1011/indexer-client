package types

type Token struct {
	Address      string `json:"address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Decimals     uint8  `json:"decimals"`
	Creator      string `json:"creator"`
	CreatedAt    int64  `json:"created_at"`
	CreationHash string `json:"creation_hash"`
	ChainID      int16  `json:"chain_id"`
}
