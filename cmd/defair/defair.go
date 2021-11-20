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
			{
				Name: "balance",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "network",
						Aliases: []string{"n"},
						Value:   "all",
					},
				},
				Action: commands.Balance,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
