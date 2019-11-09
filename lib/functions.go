package lib

import (
	"log"
	"strconv"
)

func StringContains(data []string, s string) bool {
	for _, d := range data {
		if d == s {
			return true
		}
	}
	return false
}

func String2Float(s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("Command failuer at String2Float: args: %s", s)
	}
	return num
}
