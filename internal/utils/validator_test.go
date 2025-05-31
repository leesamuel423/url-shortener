package utils

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{"Valid HTTP URL", "http://example.com", true},
		{"Valid HTTPS URL", "https://example.com", true},
		{"Valid URL with path", "https://example.com/path/to/page", true},
		{"Valid URL with query", "https://example.com?q=test", true},
		{"Invalid - no scheme", "example.com", false},
		{"Invalid - empty", "", false},
		{"Invalid - FTP scheme", "ftp://example.com", false},
		{"Invalid - no host", "https://", false},
		{"Invalid - malformed", "ht!tp://example", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidURL(tt.url)
			if result != tt.expected {
				t.Errorf("IsValidURL(%s) = %v, want %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestIsSameDomain(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		domain   string
		expected bool
	}{
		{"Same domain", "https://example.com/path", "https://example.com", true},
		{"Same domain with www", "https://www.example.com", "https://www.example.com", true},
		{"Different domains", "https://other.com", "https://example.com", false},
		{"Different schemes", "http://example.com", "https://example.com", true},
		{"Case insensitive", "https://EXAMPLE.com", "https://example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameDomain(tt.url, tt.domain)
			if result != tt.expected {
				t.Errorf("IsSameDomain(%s, %s) = %v, want %v", tt.url, tt.domain, result, tt.expected)
			}
		})
	}
}
