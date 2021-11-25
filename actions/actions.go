package actions

import (
	"github.com/ahstn/defair/platform"
)

// LiquidityPools uses platform.Chain to retrieve the Pools a Wallet is
// providing Liquidity for across multiple networks & protocols.
func LiquidityPools(address string) error {
	y := platform.YamlIndex{Path: "config.yaml"}
	c, err := y.Read()
	if err != nil {
		return err
	}

	e := platform.EthClient{}
	err = e.LiquidityPools(address, c.Networks["avalanche"].Endpoint, c.Networks["avalanche"].Markets)

	return nil
}
