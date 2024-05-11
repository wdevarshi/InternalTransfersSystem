package validator

import "testing"

func contains(s string, c string) bool {
	for _, char := range s {
		if string(char) == c {
			return true
		}
	}
	return false
}

func TestContainsSpecialCharacters(t *testing.T) {
	// Test case with special characters
	if !containsSpecialCharacters("abc!def") {
		t.Errorf("Expected special characters to be found, but not found")
	}

	// Test case with special characters
	if !containsSpecialCharacters("abc def") {
		t.Errorf("Expected special characters to be found, but not found")
	}

	// Test case without special characters
	if containsSpecialCharacters("abcdefgh") {
		t.Errorf("Expected no special characters to be found, but found")
	}
}
