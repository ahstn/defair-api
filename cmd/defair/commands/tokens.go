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

// Tokens defines the necessary flags and action for retrieving a wallet's tokens.
var Tokens = cli.Command{
	Name: "tokens",
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

		tokens, err := actions.Tokens(c.Args().Get(0), f, y, e)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range tokens {
			fmt.Printf("token=[%s - %q], Balance=[%.2f]\n", t.Symbol, t.Name, t.Balance)
		}
		return nil
	},
}
