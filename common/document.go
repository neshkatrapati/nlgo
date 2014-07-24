package common

import (
	"strings"
)

type Document struct {
	Sentences []*Sentence
}
type Sentence struct {
	Tokens []*string
}

func NewDocument(RawText string) *Document {
	sentences := strings.Split(RawText, "\n")
	return &Document{Sentences: &sentences}
}
