package util

import "testing"

func TestLemmatize(t *testing.T) {
	result := Lemmatize("jumps")
	expected := "jump"
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}

func TestLemmatizeAlreadyBaseForm(t *testing.T) {
	result := Lemmatize("jump")
	expected := "jump"
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}

func TestLemmatizeUnknownWord(t *testing.T) {
	result := Lemmatize("foobar")
	expected := "foobar"
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}

func TestLemmatizeUppercase(t *testing.T) {
	result := Lemmatize("JUMPS")
	expected := "jump"
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}

func TestLemmatizeMultiplePossibleLemmas(t *testing.T) {
	result := Lemmatize("swim")
	expected := "swim"
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}

func TestLemmatizeEmptyString(t *testing.T) {
	result := Lemmatize("")
	expected := ""
	if result != expected {
		t.Errorf("Test case failed: got %s, want %s", result, expected)
	}
}
