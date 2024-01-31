package client

import (
	"context"
	"testing"

	"github.com/AR1011/indexer-client/types"
)

func TestGetBlockNumber(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	result, err := c.GetBlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal(errNilResult)
	}
}

func TestGetChains(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	result, err := c.GetChains(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal(errNilResult)
	}
}

func TestGetBlockTimestamps(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	params := types.GetBlockTimestampsParams{
		ChainID:   1,
		FromBlock: 1,
		ToBlock:   10,
	}

	result, err := c.GetBlockTimestamps(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal(errNilResult)
	}
}

func TestGetBlockAtTimestamp(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	params := types.GetBlockAtTimestampParams{
		ChainID:   1,
		Timestamp: 1680911893,
	}

	result, err := c.GetBlockAtTimestamp(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal(errNilResult)
	}
}

func TestGetHeights(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	result, err := c.GetHeight(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal(errNilResult)
	}
}

func TestGetPairCount(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	result, err := c.GetPairCount(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	if result == 0 {
		t.Fatal("no pairs")
	}
}

func TestGetTokenCount(t *testing.T) {
	t.Parallel()

	c := New("http://localhost:8080")
	c.WithDebug()

	ctx := context.Background()

	result, err := c.GetTokenCount(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	if result == 0 {
		t.Fatal("no tokens")
	}
}
