package main

import (
	"math/rand"
	"strings"
)

type Quote struct {
	quote  string
	author string
}

func searchQuote(heystack []Quote, needle string) []Quote {
	res := []Quote{}
	for _, val := range heystack {
		for _, word := range strings.Split(val.quote, " ") {
			if needle == word {
				res = append(res, val)
			}
		}
	}
	return res
}

func getQuote(from []Quote) *Quote {
	res := rand.Intn(len(from))
	return &from[res]
}

func main() {

}
