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
	networks := []string{"avalanche", "harmony", "aurora"}
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


