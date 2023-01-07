package util

import (
	"log"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

var lem *golem.Lemmatizer

func Lemmatize(word string) string {
	if lem == nil {
		lemmatizer, err := golem.New(en.New())
		if err != nil {
			log.Print("Lemmatizer not working: ", err)
			return word
		}
		lem = lemmatizer
	}
	return lem.Lemma(word)
}
