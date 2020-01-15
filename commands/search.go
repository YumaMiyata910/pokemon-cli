package commands

import (
	"github.com/ymiyata910/pokemon-cli/action"
	"github.com/urfave/cli"
)

func setSearch() cli.Command {
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
