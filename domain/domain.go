package domain

import "time"

// Wallet holds the user's wallet entity
type Wallet struct {
	// Address to the wallet (i.e. 0x123...)
	Address string
	// Networks the wallet is configured for
	Networks []string
}

// TokenList represents the traditional JSON structure used by ERC-20 token lists - tokenlists.org
type TokenList struct {
	Name     string   `json:"name"`
	LogoURI  string   `json:"logoURI"`
	Keywords []string `json:"keywords"`
	Version  struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Tokens    []Token   `json:"tokens"`
}

// TokenPair represents a pair of ERC20 Balances.
type TokenPair struct {
	Token0 Token
	Token1 Token
}

// Token represents an ERC20 Token's address & symbol.
type Token struct {
	Address  string  `json:"address"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Balance  float64 `json:"balance"`
	Decimals int     `json:"decimals"`
	LogoURI  string  `json:"logoURI,omitempty"`
}

// LiquidityPool represents a pool existing of two ERC20 Balances to provide liquidity for.
type LiquidityPool struct {
	// Address is the EVM address of the LP's smart contract.
	Address string `json:"address,omitempty"`
	// Balance is the user's balance for the LP.
	Balance float64 `json:"balance,omitempty"`
	// Rewards is the user's outstanding reward amount.
	Rewards float64 `json:"rewards,omitempty"`
	// Pair is the LP's underlying TokenPair / assets.
	Pair TokenPair `json:"pair,omitempty"`
	// Market is the Market which the LP belongs to.
	Market Market `json:"market,omitempty"`
}

// DataFilter is used to filter operations throughout the API.
// i.e. Get all LPs from a specific Network.
type DataFilter struct {
	Networks  []string
	Protocols []string
}

// Index is a directory of on-chain contract and token Addresses.
type Index struct {
	Networks map[string]Network `yaml:"networks"`
}

// Network represents different data for a specific EVM network, including it's RPC API Endpoint.
type Network struct {
	Endpoint  string       `yaml:"endpoint"`
	Tokens    TokenSources `yaml:"tokens,omitempty"`
	Exchanges []Exchange   `yaml:"exchanges,omitempty"`
	Markets   []Market     `yaml:"markets,omitempty"`
}

// TokenSources represents sources for on-chain token addresses.
type TokenSources struct {
	Lists      []string `yaml:"lists"`
	Additional []string `yaml:"additional"`
}

// Exchange represents an on-chain Exchange, it's Governance Token and related Smart Contracts.
type Exchange struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

// Market represents an on-chain liquidity provider, it's Governance Token and related Smart Contracts.
type Market struct {
	Name  string   `yaml:"name" json:"name"`
	Token string   `yaml:"token"`
	Chef  []string `yaml:"chef"`
}
