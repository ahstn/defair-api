package actions

import (
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
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

	networks := networkFilter(filter)

	for _, network := range networks {
		networkPools, err := e.LiquidityPools(address, c.Networks[network])
		if err != nil {
			return pools, nil
		}
		pools = append(pools, networkPools...)
	}

	return pools, nil
}
