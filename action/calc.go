package action

import (
	"fmt"
	"log"
	"math"

	"github.com/YumaMiyata910/pokemon-cli/constant"
	"github.com/YumaMiyata910/pokemon-cli/lib"
	"github.com/YumaMiyata910/pokemon-cli/pokemon"

	"github.com/urfave/cli"
)

func IVCalc(c *cli.Context) error {
	if c.NArg() < 5 {
		return fmt.Errorf("There are not enough arguments. Require 5, but take %d", c.NArg())
	}
	name := c.Args().Get(0)
	charStr := c.Args().Get(1)
	realStr := c.Args().Get(2)
	effortStr := c.Args().Get(3)
	levelStr := c.Args().Get(4)

	poke := pokemon.GetByName(name)

	var baseInt int
	var target constant.StatusType
	var output string

	t := c.String("target")
	switch {
	case lib.StringContains([]string{"H", "h"}, t):
		target = constant.HP
		baseInt = poke.HitPoint
	case lib.StringContains([]string{"A", "a"}, t):
		target = constant.Attack
		baseInt = poke.Attack
	case lib.StringContains([]string{"B", "b"}, t):
		target = constant.Defence
		baseInt = poke.Defense
	case lib.StringContains([]string{"C", "c"}, t):
		target = constant.SpecialAttack
		baseInt = poke.SpecialAttack
	case lib.StringContains([]string{"D", "d"}, t):
		target = constant.SpecialDefence
		baseInt = poke.SpecialDefense
	case lib.StringContains([]string{"S", "s"}, t):
		target = constant.Speed
		baseInt = poke.Speed
	default:
		log.Fatalf("Command failuer. Require target flag {H|A|B|C|D|S}. But specified %s", t)
	}

	baseVal := float64(baseInt)
	effortVal := lib.String2Float(effortStr)
	realVal := lib.String2Float(realStr)
	level := lib.String2Float(levelStr)
	charVal := constant.GetCollectionVal(charStr, target)
	if target == "HP" {
		output = calcIvHP(baseVal, effortVal, realVal, level)
	} else {
		output = calcIv(baseVal, effortVal, realVal, level, charVal)
	}

	fmt.Printf("%s: %s\n", target, output)

	return nil
}

func calcIvHP(b, e, r, l float64) string {
	i := (100*r-1000)/l - (2*b + e/4 + 100)
	return fmt.Sprintf("%.0f - %.0f", math.Floor(i), math.Ceil(i))
}

func calcIv(b, e, r, l, c float64) string {
	i := ((100 * r) / (c * l)) - (500 / l) - (2*b + e/4)
	return fmt.Sprintf("%.0f - %.0f", math.Floor(i), math.Ceil(i))
}
