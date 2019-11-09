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
				Usage:  "calc pokemon individual value. Args: `{name} {character} {real_value} {effort_value} {level}`",
				Action: action.IVCalc,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "target, t",
						Usage: "Target calc iv target in `{H|A|B|C|D|S}`",
					},
				},
			},
		},
	}
}
