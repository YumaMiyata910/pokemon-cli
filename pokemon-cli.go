package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/ymiyata910/pokemon-cli/commands"
	"github.com/ymiyata910/pokemon-cli/pokemon"
)

func init() {
	pokemon.Load()
}

func main() {
	app := cli.NewApp()
	app.Name = "pokemon cli"
	app.Usage = "We can search pokemon info from data"
	app.Commands = commands.SetCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
