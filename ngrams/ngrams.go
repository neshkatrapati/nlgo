package ngrams

import (
	"strconv"
	"strings"
)

type NGram struct {
	Tokens []string
	Count  int
}
type NGrams struct {
	N     int
	Grams map[string]*NGram
}

func (ngram *NGram) inc(Count int) {
	ngram.Count += Count
}
func (ngram *NGram) String() string {
	return strings.Join(ngram.Tokens, " ") + " - " + strconv.Itoa(ngram.Count)
}
func NewNGrams(count int) *NGrams {
	ngrams := &NGrams{N: count, Grams: make(map[string]*NGram)}
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
			if _, ok := ngrams.Grams[realtoken]; ok {
				ngrams.Grams[realtoken].Count++

			} else {
				ngram := NewNGram(realtoken, 1)
				ngrams.Grams[realtoken] = ngram
			}
		}
	}
}
