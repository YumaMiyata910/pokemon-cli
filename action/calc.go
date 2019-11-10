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

func RVCalc(c *cli.Context) error {
	if c.NArg() < 5 {
		return fmt.Errorf("There are not enough arguments. Require 5, but take %d", c.NArg())
	}
	name := c.Args().Get(0)
	charStr := c.Args().Get(1)
	ivStr := c.Args().Get(2)
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
	ivVal := lib.String2Float(ivStr)
	level := lib.String2Float(levelStr)
	charVal := constant.GetCollectionVal(charStr, target)
	if target == "HP" {
		output = calcRvHP(baseVal, effortVal, ivVal, level)
	} else {
		output = calcRv(baseVal, effortVal, ivVal, level, charVal)
	}

	fmt.Printf("%s: %s\n", target, output)

	return nil
}

func calcIvHP(b, e, r, l float64) string {
	r1 := r
	r2 := r + 1 - 0.000001
	i1 := (100*r1-1000)/l - (2*b + e/4 + 100)
	i2 := (100*r2-1000)/l - (2*b + e/4 + 100)
	if math.Floor(i1) < 0 {
		i1 = 0
	}
	if math.Ceil(i2) > 31 {
		i2 = 31
	}
	return fmt.Sprintf("%.0f - %.0f", math.Ceil(i1), math.Floor(i2))
}

func calcIv(b, e, r, l, c float64) string {
	r1 := r
	r2 := r + 1 - 0.000001
	i1 := ((100 * r1) / (c * l)) - (500 / l) - (2*b + e/4)
	i2 := ((100 * r2) / (c * l)) - (500 / l) - (2*b + e/4)
	if math.Floor(i1) < 0 {
		i1 = 0
	}
	if math.Ceil(i2) > 31 {
		i2 = 31
	}
	return fmt.Sprintf("%.0f - %.0f", math.Ceil(i1), math.Floor(i2))
}

func calcRvHP(b, e, i, l float64) string {
	r := math.Floor((b*2+i+e/4))*(l/100) + 10 + l
	return fmt.Sprintf("%.0f", math.Floor(r))
}

func calcRv(b, e, i, l, c float64) string {
	r := math.Floor((math.Floor((b*2+i+e/4))*(l/100) + 5)) * c
	return fmt.Sprintf("%.0f", math.Floor(r))
}
