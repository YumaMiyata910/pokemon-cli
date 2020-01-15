package subcommands

import (
	"github.com/urfave/cli"
	"github.com/ymiyata910/pokemon-cli/action"
)

func SetCalc() cli.Command {
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
			{
				Name:   "rv",
				Usage:  "calc pokemon real value from individual value. Args `{name} {character} {individual_value} {effort_value} {level}`",
				Action: action.RVCalc,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "target, t",
						Usage: "Target calc rv target in `{H|A|B|C|D|S}`",
					},
				},
			},
		},
	}
}
