package constant

type Character struct {
	Name       string
	UpStatus   StatusType
	DownStatus StatusType
}

var Chars []Character = []Character{
	{"がんばりや", "", ""},
	{"さみしがり", Attack, Defence},
	{"いじっぱり", Attack, SpecialAttack},
	{"やんちゃ", Attack, SpecialDefence},
	{"ゆうかん", Attack, Speed},
	{"ずぶとい", Defence, Attack},
	{"すなお", "", ""},
	{"わんぱく", Defence, SpecialAttack},
	{"のうてんき", Defence, SpecialDefence},
	{"のんき", Defence, Speed},
	{"ひかえめ", SpecialAttack, Attack},
	{"おっとり", SpecialAttack, Defence},
	{"てれや", "", ""},
	{"うっかりや", SpecialAttack, SpecialDefence},
	{"れいせい", SpecialAttack, Speed},
	{"おだやか", SpecialDefence, Attack},
	{"おとなしい", SpecialDefence, Defence},
	{"しんちょう", SpecialDefence, SpecialAttack},
	{"きまぐれ", "", ""},
	{"なまいき", SpecialDefence, Speed},
	{"おくびょう", Speed, Attack},
	{"せっかち", Speed, Defence},
	{"ようき", Speed, SpecialAttack},
	{"むじゃき", Speed, SpecialDefence},
	{"まじめ", "", ""},
}

func GetByName(name string) Character {
	for _, char := range Chars {
		if name == char.Name {
			return char
		}
	}
	return Character{}
}

func GetCollectionVal(name string, target StatusType) float64 {
	char := GetByName(name)
	switch {
	case char.UpStatus == target:
		return 1.1
	case char.DownStatus == target:
		return 0.9
	default:
		return 1.0
	}
}
