package platform

import "strings"

type Token struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Amount float64
}

type TokenPair struct {
	Reserve0    string `json:"reserve0"`
	Reserve1    string `json:"reserve1"`
	TotalSupply string `json:"totalSupply"`
	Token0      *Token `json:"token0"`
	Token1      Token  `json:"token1"`
}

type Positions struct {
	ID      string    `json:"id"`
	Pair    TokenPair `json:"pair"`
	Balance float64
}

type PoolsResponse struct {
	Pools []Positions `json:"pools"`
}

type SnapshotsResponse struct {
	Snapshots []struct {
		Balance   string    `json:"balance"`
		Block     int       `json:"block"`
		ID        string    `json:"id"`
		Pool      Positions `json:"pool"`
		Timestamp int       `json:"timestamp"`
		Wallet    struct {
			ID string `json:"id"`
		} `json:"wallet"`
	} `json:"snapshots"`
}

type MintsResponse struct {
	Mints []struct {
		Amount0     string    `json:"amount0"`
		Amount1     string    `json:"amount1"`
		Pair        TokenPair `json:"pair"`
		Sender      string    `json:"sender"`
		Timestamp   string    `json:"timestamp"`
		To          string    `json:"to"`
		Transaction struct {
			BlockNumber string `json:"blockNumber"`
			Timestamp   string `json:"timestamp"`
		} `json:"transaction"`
	} `json:"mints"`
}

type SwapsResponse struct {
	Users []struct {
		ID                 string      `json:"id"`
		LiquidityPositions []Positions `json:"liquidityPositions"`
	} `json:"users"`
}

var (
	snapshotsQuery = strings.Join([]string{tokenFragment, `
		query ($addr: String!) {
			snapshots: liquidityPositionSnapshots(
				orderBy: timestamp,
				orderDirection: desc,
				where: {user: $addr, liquidityTokenBalance_gt: 0}
			) {
				id
				timestamp
				block
				wallet: user {
					id
				}
				pool: liquidityPosition {
					id
					pair {
						...PairInfo
					}
				}
				balance: liquidityTokenBalance
			}
		}
	`}, "\n")

	swapsQuery = strings.Join([]string{tokenFragment, `
		query ($addr: String!) {
			users(where: {id: $addr}) {
				id
				pool: liquidityPositions {
					id
					pair {
						...PairInfo
					}
				}
			}
		}
	`}, "\n")
)

const (
	tokenFragment = `
		fragment TokenInfo on Token {
			symbol
			name
		}
		fragment PairInfo on Pair {
			totalSupply
			reserve0
			reserve1
			token0 {
				...TokenInfo
			}
			token1 {
				...TokenInfo
			}
		}
	`

	poolsQuery = `
		query ($addr: String!) {
			pools: users(where: {id: $addr}) {
				id
				liquidityPositions {
					id
					pair {
						...PairInfo
					}
				}
			}
		}
	`

	mintsQuery = `
		query ($addr: String!) {
			mints: mints(orderDirection: asc, where: {to: $addr}) {
				sender
				to
				timestamp
				transaction {
					timestamp
					blockNumber
				}
				amount0
				amount1
				pair {
					...PairInfo
				}
			}
		}
	`
)
