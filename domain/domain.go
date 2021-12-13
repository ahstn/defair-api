package domain

// Wallet holds the user's wallet entity
type Wallet struct {
	// Address to the wallet (i.e. 0x123...)
	Address string
	// Networks the wallet is configured for
	Networks []string
}

// TokenPair represents a pair of ERC20 Tokens.
type TokenPair struct {
	Token0 Token
	Token1 Token
}

// Token represents an ERC20 Token's address & symbol.
type Token struct {
	Address string
	Symbol string
}

// LiquidityPool represents a pool existing of two ERC20 Tokens to provide liquidity for.
type LiquidityPool struct {
	// Address is the EVM address of the LP's smart contract.
	Address string    `json:"address,omitempty"`
	// Balance is the user's balance for the LP.
	Balance float32   `json:"balance,omitempty"`
	// Rewards is the user's outstanding reward amount.
	Rewards float32   `json:"rewards,omitempty"`
	// Pair is the LP's underlying TokenPair / assets.
	Pair    TokenPair `json:"pair,omitempty"`
	// Market is the Market which the LP belongs to.
	Market  Market    `json:"market,omitempty"`
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
	Endpoint  string            `yaml:"endpoint"`
	Tokens    map[string]string `yaml:"tokens,omitempty"`
	Exchanges []Exchange        `yaml:"exchanges,omitempty"`
	Markets   []Market          `yaml:"markets,omitempty"`
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
