package utils

import (
	"strings"
	"testing"
)

func TestGenerateShortCode(t *testing.T) {
	// Test length
	code := GenerateShortCode()
	if len(code) != shortCodeLength {
		t.Errorf("Expected length %d, got %d", shortCodeLength, len(code))
	}

	// Test uniqueness
	codes := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		code := GenerateShortCode()
		if codes[code] {
			t.Errorf("Duplicate code generated: %s", code)
		}
		codes[code] = true
	}

	// Test character set
	for _, char := range code {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("Invalid character in code: %c", char)
		}
	}
}
