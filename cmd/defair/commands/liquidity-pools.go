package commands

import (
	"log"

	"github.com/ahstn/defair/actions"
	"github.com/urfave/cli/v2"
)

// LiquidityPools defines the necessary flags and action for retrieving the
// Pools a Wallet is providing Liquidity for.
var LiquidityPools = cli.Command{
	Name: "liquidity-pools",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "network",
			Aliases: []string{"n"},
			Value:   "all",
		},
	},
	Action: func(c *cli.Context) error {
		err := actions.LiquidityPools()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
