package main

import (
	"fmt"
	"github.com/neshkatrapati/nlgo/common"
	"github.com/neshkatrapati/nlgo/ngrams"
	"os"
	"sort"
	"strconv"
)

func Wait() {
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func main() {

	var NS string = os.Args[1] // Give N as command line arg
	N, _ := strconv.Atoi(NS)
	ng := ngrams.NewNGrams(N)
	var String string = "I am sam sam I am I do not eat ham"

	fmt.Println("String Given:", String)
	ng.FromString(String, " ")

	for _, ngram := range ng.Grams {
		fmt.Println(ngram)
	}

	ngramlist := ng.AsList()
	sort.Sort(ngramlist)
	fmt.Println("Sorted ::")

	for _, ngram := range ngramlist {
		fmt.Println(ngram)
	}

	fmt.Println("Min ::")
	fmt.Println(ng.Min())

	fmt.Println("Max ::")
	fmt.Println(ng.Max())

}
