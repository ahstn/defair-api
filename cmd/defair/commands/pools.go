package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"

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
		y := platform.YamlIndex{Path: "./config.yaml"}
		e := platform.EthClient{}
		f := domain.DataFilter{Networks: strings.Split(c.String("network"), ",")}

		pools, err := actions.LiquidityPools(c.Args().Get(0), f, y, e)
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range pools {
			fmt.Printf("Market=[%s], LP=[%s], Balance=[%v], Rewards=[%v]\n",
				p.Market.Name, p.Address, p.Balance, p.Rewards)
		}
		return nil
	},
}
