package commands

import (
	"fmt"

	"github.com/ahstn/defair/cmd/defair/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

// Server defines the necessary flags and action for initialising a Gin Server.
var Server = cli.Command{
	Name: "server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "9000",
		},
	},
	Action: func(c *cli.Context) error {
		r := gin.Default()
		r.GET("/wallet/:address", endpoints.Balance)
		r.Run(fmt.Sprintf(":%s", c.String("port")))
		return nil
	},
}
