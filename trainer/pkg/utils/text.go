package util

import (
	"strings"
	"unicode"
)

//Tokenize and clean text
func Tokenize(texts []string, stopWords map[string]struct{}) []string {
	result := make([]string, 0)

	for _, t := range texts {
		text := strings.ToLower(t)
		tokens := strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) })
		withoutStopWords := cleanTokenizedText(tokens, stopWords)
		result = append(result, withoutStopWords...)
	}

	return removeDuplicates(result)
}

//Removes stop words from string
func cleanTokenizedText(tokens []string, stopWords map[string]struct{}) []string {
	cleaned := []string{}

	for _, word := range tokens {
		if _, ok := stopWords[strings.ToLower(word)]; !ok {
			lemmatized := Lemmatize(word)
			cleaned = append(cleaned, lemmatized)
		}
	}
	return cleaned
}

//Remove duplicates from the strings
func removeDuplicates(text []string) []string {
	uniqueStrings := make([]string, 0, len(text))
	seen := make(map[string]struct{})

	for _, s := range text {
		if _, ok := seen[s]; !ok {
			uniqueStrings = append(uniqueStrings, s)
			seen[s] = struct{}{}
		}
	}
	return uniqueStrings
}
