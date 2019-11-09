package main

import (
	"testing"

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

	for _, d := range testData {
		c := testcli.Command("pokemon-cli", "search", "name", d)
		c.Run()
		if !c.Success() {
			t.Fatalf("Expected to succeed, but failed: %s", c.Error())
		}

		if !c.StdoutContains(d) {
			t.Fatalf("Expected %q to contain %q", c.Stdout(), d)
		}
	}
}
