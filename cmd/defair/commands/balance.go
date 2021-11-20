package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Balance defines the necessary flags and action for retrieving a Wallet balance(s).
var Balance = cli.Command{
	Name: "balance",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "network",
			Aliases: []string{"n"},
			Value:   "all",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Printf("Checking balance for %q on network %q", c.Args().Get(0), c.String("network"))
		return nil
	},
}
