package actions

import (
	"fmt"

	"github.com/ahstn/defair/platform"
)

// LiquidityPools uses platform.Chain to retrieve the Pools a Wallet is
// providing Liquidity for across multiple networks & protocols.
func LiquidityPools() error {
	y := platform.YamlIndex{Path: "config.yaml"}
	c, err := y.Read()
	if err != nil {
		return err
	}

	fmt.Println("Avax Tokens")
	for k, v := range c.Networks["avalanche"].Tokens {
		fmt.Printf("Avalanche Token: %q - %q\n", k, v)
	}

	return nil
}
