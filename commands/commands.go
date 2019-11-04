package commands

import (
	"github.com/urfave/cli"
)

func SetCommands() []cli.Command {
	var commands []cli.Command

	commands = append(commands, setSearch())
	commands = append(commands, setCalc())

	return commands
}
