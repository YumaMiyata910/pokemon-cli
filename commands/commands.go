package commands

import (
	"github.com/urfave/cli"
	"github.com/ymiyata910/pokemon-cli/subcommands"
)

func SetCommands() []cli.Command {
	var commands []cli.Command

	commands = append(commands, subcommands.SetSearch())
	commands = append(commands, subcommands.SetCalc())

	return commands
}
