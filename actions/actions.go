package actions

import (
	"fmt"
	"github.com/ahstn/defair/platform"
	"log"
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
	lps, err := e.LiquidityPools(address, c.Networks["avalanche"].Endpoint, c.Networks["avalanche"].Markets)
	if err != nil {
		log.Fatal(err)
	}

	for _, lp := range lps {
		fmt.Printf("LP (%s), Balance=[%v], Rewards=[%v]", lp.Address, lp.Balance, lp.Rewards)
	}

	return nil
}
