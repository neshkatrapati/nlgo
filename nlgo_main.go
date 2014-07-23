package main

import (
	"fmt"
	"github.com/neshkatrapati/nlgo/ngrams"
)

func main() {
	ng := ngrams.NewNGrams(2)
	var String string = "I am sam sam I am I do not eat ham"
	ng.FromString(String, " ")
	fmt.Println(String)
	for _, ngram := range ng.Grams {
		fmt.Println(ngram)
	}
}
