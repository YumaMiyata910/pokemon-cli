package main_test

import (
	"testing"

	"github.com/ymiyata910/pokemon-cli/constant"
	"github.com/rendon/testcli"
)

func TestSearchName(t *testing.T) {
	testData := []string{
		"フシギダネ",
		"チコリータ",
		"キモリ",
		"ナエトル",
		"ツタージャ",
		"ハリマロン",
		"モクロウ",
		"マーシャドー",
	}

	for i, d := range testData {
		c := testcli.Command("pokemon-cli", "search", "name", d)
		c.Run()
		if !c.Success() {
			t.Fatalf("case:【%d】 Expected to succeed, but failed: %s", i, c.Error())
		}

		if !c.StdoutContains(d) {
			t.Fatalf("case:【%d】 Expected %q to contain %q", i, c.Stdout(), d)
		}
	}
}

func TestCalcIv(t *testing.T) {
	testData := []struct {
		name       string
		target     string
		char       string
		realVal    string
		effortVal  string
		level      string
		statusType constant.StatusType
		iv         string
	}{
		{"ハピナス", "H", "おだやか", "362", "252", "50", constant.HP, "31"},
		{"ミュウ", "A", "いじっぱり", "167", "252", "50", constant.Attack, "31"},
		{"ミュウ", "A", "いじっぱり", "328", "252", "100", constant.Attack, "31"},
		{"ミュウ", "A", "わんぱく", "152", "252", "50", constant.Attack, "31"},
		{"ミュウ", "C", "いじっぱり", "108", "0", "50", constant.SpecialAttack, "31"},
		{"ミュウ", "C", "いじっぱり", "100", "0", "50", constant.SpecialAttack, "14"},
		{"ミュウ", "D", "いじっぱり", "178", "120", "70", constant.SpecialDefence, "18"},
		{"ミュウ", "D", "ようき", "105", "0", "50", constant.SpecialDefence, "0"},
		{"ミュウ", "S", "ようき", "149", "252", "50", constant.Speed, "0"},
		{"ミュウ", "S", "おくびょう", "328", "252", "100", constant.Speed, "31"},
	}

	for i, d := range testData {
		c := testcli.Command("pokemon-cli", "calc", "iv", "--target", d.target, d.name, d.char, d.realVal, d.effortVal, d.level)
		c.Run()
		if !c.Success() {
			t.Fatalf("case:【%d】 Expected to succeed, but failed: %s", i, c.Error())
		}

		if !c.StdoutContains(string(d.statusType)) {
			t.Fatalf("case:【%d】 Expected %q to contain %q", i, c.Stdout(), d.statusType)
		}

		if !c.StdoutContains(d.iv) {
			t.Fatalf("case:【%d】 Expected %q to contain %q", i, c.Stdout(), d.iv)
		}
	}
}

func TestCalcRv(t *testing.T) {
	testData := []struct {
		name       string
		target     string
		char       string
		ivVal      string
		effortVal  string
		level      string
		statusType constant.StatusType
		rv         string
	}{
		{"ハピナス", "H", "おだやか", "31", "252", "50", constant.HP, "362"},
		{"ミュウ", "A", "いじっぱり", "31", "252", "50", constant.Attack, "167"},
		{"ミュウ", "A", "いじっぱり", "31", "252", "100", constant.Attack, "328"},
		{"ミュウ", "A", "わんぱく", "31", "252", "50", constant.Attack, "152"},
		{"ミュウ", "C", "いじっぱり", "31", "0", "50", constant.SpecialAttack, "108"},
		{"ミュウ", "C", "いじっぱり", "14", "0", "50", constant.SpecialAttack, "100"},
		{"ミュウ", "D", "いじっぱり", "18", "120", "70", constant.SpecialDefence, "178"},
		{"ミュウ", "D", "ようき", "0", "0", "50", constant.SpecialDefence, "105"},
		{"ミュウ", "S", "ようき", "0", "252", "50", constant.Speed, "149"},
		{"ミュウ", "S", "おくびょう", "31", "252", "100", constant.Speed, "328"},
	}

	for i, d := range testData {
		c := testcli.Command("pokemon-cli", "calc", "rv", "--target", d.target, d.name, d.char, d.ivVal, d.effortVal, d.level)
		c.Run()
		if !c.Success() {
			t.Fatalf("case:【%d】 Expected to succeed, but failed: %s", i, c.Error())
		}

		if !c.StdoutContains(string(d.statusType)) {
			t.Fatalf("case:【%d】 Expected %q to contain %q", i, c.Stdout(), d.statusType)
		}

		if !c.StdoutContains(d.rv) {
			t.Fatalf("case:【%d】 Expected %q to contain %q", i, c.Stdout(), d.rv)
		}
	}
}
