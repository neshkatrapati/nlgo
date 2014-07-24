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

type NGramList []*NGram

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

func (ngrams *NGrams) AsList() NGramList {
	ngramlist := NGramList{}
	for _, value := range ngrams.Grams {
		ngramlist = append(ngramlist, value)
	}
	return ngramlist
}

func (ngrams *NGrams) Min() *NGram {
	var min int = 10000
	var ng *NGram
	for _, ngram := range ngrams.Grams {
		if min > ngram.Count {
			min = ngram.Count
			ng = ngram
		}
	}
	return ng
}

func (ngrams *NGrams) Max() *NGram {
	var max int = -1
	var ng *NGram
	for _, ngram := range ngrams.Grams {
		if max < ngram.Count {
			max = ngram.Count
			ng = ngram
		}
	}
	return ng
}

func (nglist NGramList) Len() int           { return len(nglist) }
func (nglist NGramList) Less(i, j int) bool { return nglist[i].Count < nglist[j].Count }
func (nglist NGramList) Swap(i, j int)      { nglist[i], nglist[j] = nglist[j], nglist[i] }

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
