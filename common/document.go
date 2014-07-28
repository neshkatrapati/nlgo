package common

import (
	"strings"
)

type Document struct {
	Sentences []*Sentence
}
type Sentence struct {
	Tokens []string
}

func NewSentence(sentence string) *Sentence {
	tokens := strings.Split(sentence, " ")
	return &Sentence{Tokens : tokens}
}
func NewDocument(RawText string) *Document {
	sentences := strings.Split(RawText, "\n")
	var sentencelist []*Sentence
	for _, sentence := range sentences {
		sent := NewSentence(sentence)
		sentencelist = append(sentencelist, sent)
	}
	return &Document{Sentences: sentencelist}
}






















