package pokemon

import (
	"fmt"
	"log"
	"os"

	"github.com/YumaMiyata910/pokemon-cli/elements"
	"github.com/gocarina/gocsv"
)

const output = `名前: %s
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
合計: %d
`

type Pokemon struct {
	No             string           `csv:"index"`
	Name           string           `csv:"name"`
	Type1          elements.Element `csv:"-"`
	TypeName1      string           `csv:"type１"`
	Type2          elements.Element `csv:"-"`
	TypeName2      string           `csv:"type２"`
	Feature1       string           `csv:"ability１"`
	Feature2       string           `csv:"ability２"`
	Feature3       string           `csv:"hidden_ability"`
	HitPoint       int              `csv:"hp"`
	Attack         int              `csv:"attack"`
	Defense        int              `csv:"defence"`
	SpecialAttack  int              `csv:"special_attack"`
	SpecialDefense int              `csv:"special_defence"`
	Speed          int              `csv:"speed"`
	Total          int              `csv:"total"`
}

func (poke *Pokemon) FPrint() {
	fmt.Printf(
		output,
		poke.Name,
		poke.TypeName1,
		poke.TypeName2,
		poke.Feature1,
		poke.Feature2,
		poke.Feature3,
		poke.HitPoint,
		poke.Attack,
		poke.Defense,
		poke.SpecialAttack,
		poke.SpecialDefense,
		poke.Speed,
		poke.Total,
	)
}

var pokemons []*Pokemon

func Load() {
	file, err := os.Open("./pokemon/pokemon_status.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &pokemons); err != nil {
		log.Fatal(err)
	}

	for _, poke := range pokemons {
		poke.setType()
	}
}

func (poke *Pokemon) setType() {
	poke.Type1 = elements.GetElement(poke.TypeName1)
	poke.Type2 = elements.GetElement(poke.TypeName2)
}

func GetByName(name string) *Pokemon {
	for _, poke := range pokemons {
		if poke.Name == name {
			return poke
		}
	}

	return &Pokemon{}
}
