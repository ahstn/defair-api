package domain

// Wallet holds the user's wallet entity
type Wallet struct {
	// Address to the wallet (i.e. 0x123...)
	Address string
	// Networks the wallet is configured for
	Networks []string
}

type TokenPair struct {
	Token0 string
	Token1 string
}

type LiquidityPool struct {
	Address string    `json:"address,omitempty"`
	Balance float32   `json:"balance,omitempty"`
	Rewards float32   `json:"rewards,omitempty"`
	Pair    TokenPair `json:"pair,omitempty"`
	Market  Market    `json:"market,omitempty"`
}

type DataFilter struct {
	Networks  []string
	Protocols []string
}

// Index is a directory of on-chain contract and token Addresses.
type Index struct {
	Networks map[string]Network `yaml:"networks"`
}

type Network struct {
	Endpoint  string            `yaml:"endpoint"`
	Tokens    map[string]string `yaml:"tokens,omitempty"`
	Exchanges []Exchange        `yaml:"exchanges,omitempty"`
	Markets   []Market          `yaml:"markets,omitempty"`
}

type Exchange struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

type Market struct {
	Name  string   `yaml:"name" json:"name"`
	Token string   `yaml:"token"`
	Chef  []string `yaml:"chef"`
}
