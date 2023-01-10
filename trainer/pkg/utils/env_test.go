package util

import (
	"os"
	"testing"
)

func TestGetEnvVariable(t *testing.T) {
	// Test when key exists
	os.Setenv("EXAMPLE_KEY", "example_value")
	value := GetEnvVariable("EXAMPLE_KEY")
	if value != "example_value" {
		t.Errorf("Expected 'example_value', got %s", value)
	}

	// Test when key does not exist
	os.Unsetenv("EXAMPLE_KEY")
	value = GetEnvVariable("EXAMPLE_KEY")
	if value != "" {
		t.Errorf("Expected empty string, got %s", value)
	}
}
