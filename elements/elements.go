package elements

type Element struct {
	Index    int
	NameEn   string
	NameJp   string
	Critical []int
	Weak     []int
	Invalid  []int
}

var (
	Normal   = Element{1, "Normal", "ノーマル", []int{}, []int{}, []int{14}}
	Fire     = Element{2, "Fire", "ほのお", []int{5, 6, 12, 17}, []int{2, 5, 6, 12, 17, 18}, []int{}}
	Water    = Element{3, "Water", "みず", []int{2, 9, 13}, []int{2, 3, 6, 17}, []int{}}
	Electric = Element{4, "Electric", "でんき", []int{3, 10}, []int{4, 10, 17}, []int{}}
	Grass    = Element{5, "Grass", "くさ", []int{3, 9, 13}, []int{3, 4, 5, 9}, []int{}}
	Ice      = Element{6, "Ice", "こおり", []int{5, 9, 10, 15}, []int{6}, []int{}}
	Fighting = Element{7, "Fighting", "かくとう", []int{1, 6, 13, 16, 17}, []int{12, 13, 16}, []int{}}
	Poison   = Element{8, "Poison", "どく", []int{5, 18}, []int{5, 7, 8, 12, 18}, []int{}}
	Ground   = Element{9, "Ground", "じめん", []int{2, 4, 8, 13, 17}, []int{8, 13}, []int{4}}
	Flying   = Element{10, "Flying", "ひこう", []int{5, 7, 12}, []int{5, 7, 12}, []int{9}}
	Psychic  = Element{11, "Psychic", "エスパー", []int{7, 8}, []int{7, 11}, []int{}}
	Bug      = Element{12, "Bug", "むし", []int{5, 11, 16}, []int{5, 7, 9}, []int{}}
	Rock     = Element{13, "Rock", "いわ", []int{2, 6, 10, 12}, []int{1, 2, 8, 10}, []int{}}
	Ghost    = Element{14, "Ghost", "ゴースト", []int{11, 14}, []int{8, 12}, []int{1, 7}}
	Dragon   = Element{15, "Dragon", "ドラゴン", []int{15}, []int{2, 3, 4, 5}, []int{}}
	Dark     = Element{16, "Dark", "あく", []int{11, 14}, []int{14, 17}, []int{11}}
	Steel    = Element{17, "Steel", "はがね", []int{6, 13, 18}, []int{1, 5, 6, 10, 11, 12, 13, 15, 17, 18}, []int{8}}
	Fairy    = Element{18, "Fairy", "フェアリー", []int{7, 15, 16}, []int{7, 12, 17}, []int{15}}

	NoneType = Element{Index: -1}
)

var elements = []Element{
	Normal,
	Fire,
	Water,
	Electric,
	Grass,
	Ice,
	Fighting,
	Poison,
	Ground,
	Flying,
	Psychic,
	Bug,
	Rock,
	Ghost,
	Dragon,
	Dark,
	Steel,
	Fairy,
}

func GetElement(name string) Element {
	for _, ele := range elements {
		if ele.NameJp == name {
			return ele
		}
	}

	return NoneType
}
