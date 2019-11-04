package action

import (
	"fmt"

	"github.com/YumaMiyata910/pokemon-cli/pokemon"
	"github.com/urfave/cli"
)

var output string = `名前: %s
タイプ1: %s
タイプ2: %s
特性1: %s
特性2: %s
夢特性: %s

個体値
HP: %d
こうげき: %d
ぼうぎょ: %d
とくこう: %d
とくぼう: %d
すばやさ: %d
合計: %d`

func NameSearch(c *cli.Context) error {
	name := c.Args().First()
	if name == "" {
		return fmt.Errorf("Require pokemon name in argument")
	}

	poke := pokemon.GetByName(name)

	if poke.Name == "" {
		fmt.Printf("Not found pokemon.\nArgs: 【%s】\n", name)
		return nil
	}

	poke.FPrint()

	return nil
}
