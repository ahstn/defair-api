package actions

import (
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"github.com/thoas/go-funk"
)

// LiquidityPools uses platform.Chain to retrieve the Pools a Wallet is
// providing Liquidity for across multiple networks & protocols.
func LiquidityPools(
	address string, filter domain.DataFilter, y platform.Indexer, e platform.Chain,
) ([]domain.LiquidityPool, error) {
	var pools []domain.LiquidityPool

	c, err := y.Read()
	if err != nil {
		return pools, err
	}

	// if the filter isn't "all", intersect is used to ensure the Network Names provided match valid IDs.
	networks := []string{"avalanche", "harmony"}
	if len(filter.Networks) >= 1 && !funk.Contains(filter.Networks, "all") {
		networks = funk.IntersectString(networks, filter.Networks)
	}

	for _, network := range networks {
		networkPools, err := e.LiquidityPools(address, c.Networks[network])
		if err != nil {
			return pools, nil
		}
		pools = append(pools, networkPools...)
	}

	return pools, nil
}

// Tokens uses platform.Chain to retrieve a list of token's a Wallet has a positive balance for.
func Tokens(
	address string, filter domain.DataFilter, y platform.Indexer, e platform.Chain,
) ([]domain.Token, error) {
	var tokens []domain.Token

	c, err := y.Read()
	if err != nil {
		return tokens, err
	}

	// if the filter isn't "all", intersect is used to ensure the Network Names provided match valid IDs.
	networks := []string{"avalanche", "harmony"}
	if len(filter.Networks) >= 1 && !funk.Contains(filter.Networks, "all") {
		networks = funk.IntersectString(networks, filter.Networks)
	}

	for _, network := range networks {
		networkTokens, err := e.Tokens(address, c.Networks[network])
		if err != nil {
			return tokens, nil
		}
		tokens = append(tokens, networkTokens...)
	}

	return tokens, nil
}
