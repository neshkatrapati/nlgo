package main

import (
	"fmt"
	"github.com/neshkatrapati/nlgo/ngrams"
	"os"
	"strconv"
)

func main() {

	var NS string = os.Args[1] // Give N as command line arg
	N, _ := strconv.Atoi(NS)
	ng := ngrams.NewNGrams(N)
	var String string = "I am sam sam I am I do not eat ham"
	fmt.Println(String)
	ng.FromString(String, " ")

	for _, ngram := range ng.Grams {
		fmt.Println(ngram)
	}
}
