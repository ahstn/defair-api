package commands

import (
	"fmt"
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"log"
	"strings"

	"github.com/ahstn/defair/actions"
	"github.com/urfave/cli/v2"
)

// Staking defines the necessary flags and action for retrieving a wallet's staking tokens.
var Staking = cli.Command{
	Name: "staking",
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

		tokens, err := actions.Staking(c.Args().Get(0), f, y, e)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range tokens {
			fmt.Printf("token=[%s - %q], Balance=[%.3f], Ratio=[%.3f]\n", t.Symbol, t.Name, t.Balance, t.Ratio)
		}
		return nil
	},
}
