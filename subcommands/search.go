package subcommands

import (
	"github.com/urfave/cli"
	"github.com/ymiyata910/pokemon-cli/action"
)

func SetSearch() cli.Command {
	return cli.Command{
		Name:  "search",
		Usage: "search pokemon from data",
		Subcommands: []cli.Command{
			{
				Name:   "name",
				Usage:  "search from pokemon name",
				Action: action.NameSearch,
			},
		},
	}
}
