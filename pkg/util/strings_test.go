package util

import (
	"slices"
	"testing"
)

func TestDeDupeStrSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"foo"},
			expected: []string{"foo"},
		},
		{
			name:     "Duplicate elements",
			input:    []string{"foo", "foo"},
			expected: []string{"foo"},
		},
		{
			name:     "Multiple elements",
			input:    []string{"foo", "bar", "foo", "baz", "bar"},
			expected: []string{"foo", "bar", "baz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DeDupeStrSlice(tt.input)
			if len(actual) != len(tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
			if !slices.Equal(actual, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Lowercase",
			input:    "FOO",
			expected: "foo",
		},
		{
			name:     "Trim whitespace",
			input:    "  foo  ",
			expected: "foo",
		},
		{
			name:     "Replace non-alphanumeric characters",
			input:    "foo-bar_baz",
			expected: "foo-bar-baz",
		},
		{
			name:     "Remove consecutive & leading/trailing dashes",
			input:    "! Example---string  with  **multiple** non-alphanumeric---characters---! ",
			expected: "example-string-with-multiple-non-alphanumeric-characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Sanitize(tt.input)
			if actual != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
