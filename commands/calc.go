package commands

import (
	"github.com/YumaMiyata910/pokemon-cli/action"
	"github.com/urfave/cli"
)

func setCalc() cli.Command {
	return cli.Command{
		Name:  "calc",
		Usage: "calc pokemon values",
		Subcommands: []cli.Command{
			{
				Name:   "iv",
				Usage:  "calc pokemon individual value",
				Action: action.IVCalc,
			},
		},
	}
}
