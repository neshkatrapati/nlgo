package main

import (
	"fmt"
	"strconv"
	"strings"
)

type NGram struct {
	Tokens []string
	Count  int
}
type NGrams struct {
	N      int
	ngrams map[string]*NGram
}

func (ngram *NGram) inc(Count int) {
	ngram.Count += Count
}
func (ngram *NGram) String() string {
	return strings.Join(ngram.Tokens, " ") + " - " + strconv.Itoa(ngram.Count)
}
func NewNGrams(count int) *NGrams {
	ngrams := &NGrams{N: count, ngrams: make(map[string]*NGram)}
	ngrams.N = count
	return ngrams
}
func NewNGram(Token string, Count int) *NGram {
	Tokens := strings.Split(Token, ";")
	return &NGram{Tokens: Tokens, Count: Count}
}

func (ngrams *NGrams) FromString(String string, delimeter string) {
	tokens := strings.Split(String, delimeter)
	for index, _ := range tokens {
		if index+1 >= ngrams.N {
			realtoken := strings.Join(tokens[index-(ngrams.N-1):index+1], ";")
			if _, ok := ngrams.ngrams[realtoken]; ok {
				ngrams.ngrams[realtoken].Count++

			} else {
				ngram := NewNGram(realtoken, 1)
				ngrams.ngrams[realtoken] = ngram
			}
		}
	}
}

func main() {
	ngrams := NewNGrams(2)
	var String string = "I am sam sam I am I do not eat ham"
	ngrams.FromString(String, " ")
	fmt.Println(String)
	for _, ngram := range ngrams.ngrams {
		fmt.Println(ngram)
	}
}
