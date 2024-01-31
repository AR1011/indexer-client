## EVM Indexer Client

This is a client for a private EVM Indexer API. It is used to index and query data from EVM blockchains.

### Public API

The API is JSON-RPC 2.0 compliant and is served on port 8080 by default.
The available methods are:

- `idx_getBlockNumber` - Get the current block number for a given chain

- `idx_getChains` - Get the chain IDs for the supported chains

- `idx_getBlockTimestamps` - Get the timestamp for a range of block numbers

- `idx_getBlockAtTimestamp` - Get the block number at a timestamp

- `idx_findTokens` - Find tokens by using find params

- `idx_getTokenCount` - Get the total number of tokens

- `idx_findPairs` - Find pairs by using find params

- `idx_getPairCount` - Get the total number of pairs

- `idx_getWalletBalances` - Get wallet balances for an address (WIP)

- `idx_getTokenHolders` - Get token holders for a token (WIP)

- `idx_getOHLCVT` - Get OHLCV chart data for a pair (WIP)

### Private API

Private API methods require the Master API key to be set in the config file.
The available methods are:

- `auth_generateKey` - Generate a new API key

- `auth_deleteKey` - Delete an API key

- `auth_getKeyStats` - Get usage information for an API key

- `auth_getAuthMethod` - Get the current auth method

- `auth_getKeyType` - Get the type of API keys used for auth (`uuid`, `hex32`, `hex64` ...etc)

## JSON-RPC API

### inxi_getBlockNumber

Get the current block numbers for all configured chains

#### parameters: none

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getBlockNumber",
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getBlockNumber",
  "result": {
    "1": 17461068,
    "56": 32868781
  }
}
```

### idx_getChains

Get the chain info for all configured chains

#### parameters: none

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getChains",
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getChains",
  "result": [
    {
      "chain_id": 1,
      "name": "Ethereum",
      "short_name": "ETH",
      "explorer_url": "https://etherscan.io"
    },
    {
      "chain_id": 56,
      "name": "Binance Smart Chain",
      "short_name": "BSC",
      "explorer_url": "https://bscscan.com"
    }
  ]
}
```

### idx_getBlockTimestamps

Get the block timestamps for a range of block numbers

#### Parameters:

| Parameter    | Type  | Description                            |
| ------------ | ----- | -------------------------------------- |
| `chain_id`   | int64 | The blockchain network ID.             |
| `from_block` | int64 | The starting block number (inclusive). |
| `to_block`   | int64 | The ending block number (inclusive).   |

`from_block` and `to_block` are inclusive, meaning to_block and from_block will be included in the results. to_block - from_block + 1 will be the number of results.

Limit of 100,000 blocks per request.

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getBlockTimestamps",
  "params": {
    "chain_id": 1,
    "from_block": 17000000,
    "to_block": 17000010
  },
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getBlockTimestamps",
  "result": [
    {
      "block": 17000000,
      "timestamp": 1680911891
    },
    {
      "block": 17000001,
      "timestamp": 1680911903
    },
    {
      "block": 17000002,
      "timestamp": 1680911915
    }
  ]
}
```

### idx_getBlockAtTimestamp

Get the block number closest to a given timestamp

#### Parameters:

| Parameter   | Type  | Description                |
| ----------- | ----- | -------------------------- |
| `chain_id`  | int64 | The blockchain network ID. |
| `timestamp` | int64 | Timestamp                  |

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getBlockAtTimestamp",
  "params": {
    "chain_id": 1,
    "timestamp": 1680911893
  },
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getBlockAtTimestamp",
  "result": {
    "block": 17000000,
    "timestamp": 1680911891
  }
}
```

### `idx_findTokens`

Find tokens using various filters and options.

#### Parameters:

| Parameter  | Type   | Description                |
| ---------- | ------ | -------------------------- |
| `chain_id` | int64  | The blockchain network ID. |
| `filter`   | Object | Criteria to filter tokens. |
| `options`  | Object | Additional query options.  |

#### `TokenFilter` Object:

| Field        | Type   | Description                                   |
| ------------ | ------ | --------------------------------------------- |
| `address`    | string | The token's address.                          |
| `creator`    | string | The creator's address.                        |
| `from_block` | int64  | The starting block number for token creation. |
| `to_block`   | int64  | The ending block number for token creation.   |
| `name`       | string | The name of the token.                        |
| `symbol`     | string | The token's symbol.                           |
| `decimals`   | uint8  | The number of decimals for the token.         |
| `fuzzy`      | bool   | Enable fuzzy search for string fields.        |

#### `Options` Object:

| Field        | Type   | Description                            |
| ------------ | ------ | -------------------------------------- |
| `limit`      | int64  | Maximum number of results to return.   |
| `offset`     | int64  | Offset for pagination.                 |
| `sort_by`    | string | Field to sort the results by.          |
| `sort_order` | string | Order to sort the results (asc, desc). |

- Maximum limit is 10,000 tokens per request. To get more than 10,000 tokens, use the `offset` parameter to paginate through the results.
- Default sort is by `created_at` in asc order.

#### Sortable Fields:

- `address`
- `creator`
- `name`
- `symbol`
- `decimals`
- `created_at`

#### Sort Orders:

- `asc`
- `desc`

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_findTokens",
  "params": {
    "chain_id": 1,
    "filter": {
      "name": "I am not in danger",
      "symbol": "walter",
      "fuzzy": true
    },
    "options": {
      "limit": 100,
      "sort_by": "created_at",
      "sort_order": "desc"
    }
  },
  "id": "23"
}
```

#### Example Response

```json
{
  "id": "23",
  "method": "idx_findTokens",
  "result": [
    {
      "address": "0x50E7e4E7fa109A59B255bE882846f4186677f406",
      "name": "Who are you talking to right now? Who is it you think you see? Do you know how much I make a year? I mean, even if I told you, you wouldn't believe it. Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears. It ceases to exist, without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I AM the danger. A guy opens his door and gets shot, and you think that of me? No! I am the one who knocks!",
      "symbol": "WALTER",
      "decimals": 9,
      "creator": "0x8e89ac066DE630Db9658aB5FA8FeB4ae85279b30",
      "created_at": 17931545,
      "creation_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
      "chain_id": 1
    },
    {
      "address": "0x0D8da06819bC5bf57cBDC1C9E499F3B3982584Ac",
      "name": "I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot, and you think that of me? No! I am the one who knocks!",
      "symbol": "WALTER",
      "decimals": 9,
      "creator": "0x254fFf07998de67cF68e9e4CB0dC075430c01eFd",
      "created_at": 15156104,
      "creation_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
      "chain_id": 1
    }
  ]
}
```

### `idx_getTokenCount`

Get the total number of tokens for a given chain.

#### Parameters:

| Parameter  | Type  | Description                |
| ---------- | ----- | -------------------------- |
| `chain_id` | int64 | The blockchain network ID. |

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getTokenCount",
  "params": {
    "chain_id": 56
  },
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getTokenCount",
  "result": 1418709
}
```

### `idx_findPairs`

Find pairs using various filters and options.

#### Parameters:

| Parameter  | Type   | Description                |
| ---------- | ------ | -------------------------- |
| `chain_id` | int64  | The blockchain network ID. |
| `filter`   | Object | Criteria to filter pairs.  |
| `options`  | Object | Additional query options.  |

#### `PairFilter` Object:

| Field            | Type   | Description                                        |
| ---------------- | ------ | -------------------------------------------------- |
| `token0_address` | string | The address of token0                              |
| `token1_address` | string | The address of token1                              |
| `pool_address`   | string | The address of the LP                              |
| `from_block`     | int64  | The starting block                                 |
| `to_block`       | int64  | The ending block                                   |
| `fee`            | int64  | The fee of the pair (v3 only)                      |
| `tick_spacing`   | int64  | The tick spacing of the pair (v3 only)             |
| `hash`           | string | The hash of the pair                               |
| `pool_type`      | uint8  | The pool type of the pair (`2` for v2, `3` for v3) |
| `fuzzy`          | bool   | Enable fuzzy search for string fields.             |

#### `Options` Object:

| Field        | Type   | Description                            |
| ------------ | ------ | -------------------------------------- |
| `limit`      | int64  | Maximum number of results to return.   |
| `offset`     | int64  | Offset for pagination.                 |
| `sort_by`    | string | Field to sort the results by.          |
| `sort_order` | string | Order to sort the results (asc, desc). |

#### Sortable Fields:

- `token0_address`
- `token1_address`
- `pool_address`
- `fee`
- `tick_spacing`
- `hash`
- `pool_type`
- `created_at`

#### Sort Orders:

- `asc`
- `desc`

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_findPairs",
  "params": {
    "chain_id": 1,
    "filter": {
      "token1_address": "0xcf299bd11ceceeed13e0c6d155e70240de11e059"
    },
    "options": {
      "limit": 1,
      "sort_by": "created_at",
      "sort_order": "desc"
    }
  },
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_findPairs",
  "result": [
    {
      "token0_address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
      "token1_address": "0xcf299bd11ceceeed13e0c6d155e70240de11e059",
      "fee": 0,
      "tick_spacing": 0,
      "pool_address": "0xf8a8d7bbc800007b4b9325ac4938b5e0ac24002b",
      "pool_type": 2,
      "created_at": 17991353,
      "hash": "0xf2d398d34ff648c358d792e673d786c2ea0a434d27e8a316d7ba3b792cd7300c",
      "chain_id": 1
    }
  ]
}
```

### `idx_getPairCount`

Get the total number of pairs for a given chain.

#### Parameters:

| Parameter  | Type  | Description                |
| ---------- | ----- | -------------------------- |
| `chain_id` | int64 | The blockchain network ID. |

#### Example Request

```json
{
  "jsonrpc": "2.0",
  "method": "idx_getPairCount",
  "params": {
    "chain_id": 56
  },
  "id": "1"
}
```

#### Example Response

```json
{
  "id": "1",
  "method": "idx_getPairCount",
  "result": 1418709
}
```
