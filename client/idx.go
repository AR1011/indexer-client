package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/AR1011/indexer-client/types"
)

var (
	errNilResult = errors.New("result is nil")
)

func makeRPCError(err *types.JRPCError) error {
	return fmt.Errorf("rpc error. Code: %d, Error: %s", err.Code, err.Message)
}

func (c *Client) GetBlockNumber(ctx context.Context) (map[int64]int64, error) {
	var (
		result = &types.GetBlockNumberResponse{}
		params = types.GetBlockNumberRequest{}
		err    error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getBlockNumber",
		Params: params,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, err
}

func (c *Client) GetHeight(ctx context.Context, chainID int64) (*types.Heights, error) {
	var (
		result = &types.GetHeightsResponse{}
		params = types.GetHeightsRequest{
			ChainID: &chainID,
		}
		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getHeights",
		Params: params,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, err
}

func (c *Client) GetChains(ctx context.Context) ([]types.Chain, error) {
	var (
		result = &types.GetChainsResponse{}
		params = types.GetChainsRequest{}
		err    error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getChains",
		Params: params,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, err
}

func (c *Client) GetBlockTimestamps(ctx context.Context, params types.GetBlockTimestampsParams) ([]*types.BlockTimestamp, error) {
	var (
		result  = &types.GetBlockTimestampsResponse{}
		rParams = types.GetBlockTimestampsRequest{
			ChainID:   &params.ChainID,
			FromBlock: &params.FromBlock,
			ToBlock:   &params.ToBlock,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getBlockTimestamps",
		Params: rParams,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, nil

}

func (c *Client) GetBlockAtTimestamp(ctx context.Context, params types.GetBlockAtTimestampParams) (*types.BlockTimestamp, error) {
	var (
		result  = &types.GetBlockAtTimestampResponse{}
		rParams = types.GetBlockAtTimestampRequest{
			ChainID:   &params.ChainID,
			Timestamp: &params.Timestamp,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getBlockAtTimestamp",
		Params: rParams,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, nil

}

func (c *Client) FindTokens(ctx context.Context, params types.FindTokensParams) ([]*types.Token, error) {
	var (
		result  = &types.FindTokensResponse{}
		rParams = types.FindTokensRequest{
			ChainID: &params.ChainID,
			Filter:  params.Filter,
			Options: params.Options,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_findTokens",
		Params: rParams,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, nil

}

func (c *Client) GetTokenCount(ctx context.Context, chainID int64) (int64, error) {
	var (
		result  = &types.GetTokenCountResponse{}
		rParams = types.GetTokenCountRequest{
			ChainID: &chainID,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getTokenCount",
		Params: rParams,
	}, result)

	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, makeRPCError(result.Error)
	}

	return result.Result, nil

}

func (c *Client) FindPairs(ctx context.Context, params types.FindPairsParams) ([]*types.Pair, error) {
	var (
		result  = &types.FindPairsResponse{}
		rParams = types.FindPairsRequest{
			ChainID: &params.ChainID,
			Filter:  params.Filter,
			Options: params.Options,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_findPairs",
		Params: rParams,
	}, result)

	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errNilResult
	}

	if result.Error != nil {
		return nil, makeRPCError(result.Error)
	}

	return result.Result, nil

}

func (c *Client) GetPairCount(ctx context.Context, chainID int64) (int64, error) {
	var (
		result  = &types.GetPairCountResponse{}
		rParams = types.GetPairCountRequest{
			ChainID: &chainID,
		}

		err error
	)

	err = c.client.Call(ctx, c.Endpoint, &types.JRPCRequest{
		ID:     "1",
		Method: "idx_getPairCount",
		Params: rParams,
	}, result)

	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, makeRPCError(result.Error)
	}

	return result.Result, nil

}
