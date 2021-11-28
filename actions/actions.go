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
	avaxPools, err := e.LiquidityPools(address, c.Networks["avalanche"].Endpoint, c.Networks["avalanche"].Markets)
	if err != nil {
		log.Fatal(err)
	}
	onePools, err := e.LiquidityPools(address, c.Networks["harmony"].Endpoint, c.Networks["harmony"].Markets)
	if err != nil {
		log.Fatal(err)
	}

	for _, lp := range append(avaxPools, onePools...) {
		fmt.Printf("Market=[%s], LP=[%s], Balance=[%v], Rewards=[%v]\n",
			lp.Market.Name, lp.Address, lp.Balance, lp.Rewards)
	}

	return nil
}
