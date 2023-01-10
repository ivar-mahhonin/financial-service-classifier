package util

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	stopWords := map[string]struct{}{"the": {}, "is": {}}
	texts := []string{"The cat is on the mat", "The dog is in the garden"}
	result := Tokenize(texts, stopWords)
	expected := []string{"cat", "on", "mat", "dog", "in", "garden"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestTokenizeEmptyStopWords(t *testing.T) {
	stopWords := map[string]struct{}{}
	texts := []string{"The cat is on the mat", "The dog is in the garden"}
	result := Tokenize(texts, stopWords)
	expected := []string{"the", "cat", "be", "on", "mat", "dog", "in", "garden"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestCleanTokenizedText(t *testing.T) {
	stopWords := map[string]struct{}{"the": {}, "is": {}}
	tokens := []string{"The", "cat", "is", "on", "the", "mat"}
	result := cleanTokenizedText(tokens, stopWords)
	expected := []string{"cat", "on", "mat"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestTokenizeAllStopWords(t *testing.T) {
	stopWords := map[string]struct{}{"The": {}, "cat": {}, "is": {}, "on": {}, "the": {}, "mat": {}, "dog": {}, "in": {}, "garden": {}}
	texts := []string{"The cat is on the mat", "The dog is in the garden"}
	result := Tokenize(texts, stopWords)
	expected := []string{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestCleanTokenizedTextEmptyStopWords(t *testing.T) {
	stopWords := map[string]struct{}{}
	tokens := []string{"The", "cat", "is", "on", "the", "mat"}
	result := cleanTokenizedText(tokens, stopWords)
	expected := []string{"The", "cat", "be", "on", "the", "mat"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestCleanTokenizedTextAllStopWords(t *testing.T) {
	stopWords := map[string]struct{}{"The": {}, "cat": {}, "is": {}, "on": {}, "the": {}, "mat": {}}
	tokens := []string{"The", "cat", "is", "on", "the", "mat"}
	result := cleanTokenizedText(tokens, stopWords)
	expected := []string{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestCleanTokenizedTextSomeStopWords(t *testing.T) {
	stopWords := map[string]struct{}{"The": {}, "is": {}, "the": {}}
	tokens := []string{"The", "cat", "is", "on", "the", "mat"}
	result := cleanTokenizedText(tokens, stopWords)
	expected := []string{"cat", "on", "mat"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestCleanTokenizedTextEmptyTokens(t *testing.T) {
	stopWords := map[string]struct{}{"The": {}, "is": {}, "the": {}}
	tokens := []string{}
	result := cleanTokenizedText(tokens, stopWords)
	expected := []string{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	result := removeDuplicates([]string{"a", "b", "c", "a", "d", "b"})
	expected := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicatesAllDuplicates(t *testing.T) {
	result := removeDuplicates([]string{"a", "a", "a", "a", "a"})
	expected := []string{"a"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicatesEmptyList(t *testing.T) {
	result := removeDuplicates([]string{})
	expected := []string{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicatesSingleString(t *testing.T) {
	result := removeDuplicates([]string{"a"})
	expected := []string{"a"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicatesMultipleUniqueStrings(t *testing.T) {
	result := removeDuplicates([]string{"a", "b", "c", "d"})
	expected := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}
