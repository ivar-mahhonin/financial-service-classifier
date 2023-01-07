package util

import (
	"strings"
	"unicode"
)

func Tokenize(texts []string, stopWords []string) []string {
	result := make([]string, len(texts))

	for _, t := range texts {
		text := strings.ToLower(t)
		tokens := strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) })
		withoutStopWords := cleanTokenizedText(tokens, stopWords)
		result = append(result, strings.Split(withoutStopWords, " ")...)
	}

	return removeDuplicates(result)
}

//Removes stop words from string
func cleanTokenizedText(tokens []string, stopWords []string) string {
	cleaned := []string{}

	for _, word := range tokens {
		if !contains(stopWords, word) {
			lemmatized := Lemmatize(word)
			cleaned = append(cleaned, lemmatized)
		}
	}
	filteredText := strings.Join(cleaned, " ")
	return filteredText
}

//Checks if slice of strings contain string
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

//Remove duplicates from the strings
func removeDuplicates(text []string) []string {
	uniqueStrings := make(map[string]struct{})

	for _, s := range text {
		uniqueStrings[s] = struct{}{}
	}

	result := make([]string, 0, len(uniqueStrings))

	for s := range uniqueStrings {
		result = append(result, s)
	}
	return result
}
