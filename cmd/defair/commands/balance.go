package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Balance(c *cli.Context) error {
	fmt.Printf("Checking balance for %q on network %q", c.Args().Get(0), c.String("network"))
	return nil
}
