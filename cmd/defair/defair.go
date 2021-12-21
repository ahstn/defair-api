package main

import (
	"log"
	"os"

	"github.com/ahstn/defair/cmd/defair/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "defair",
		Commands: []*cli.Command{
			&commands.Balance,
			&commands.LiquidityPools,
			&commands.Tokens,
			&commands.Server,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
