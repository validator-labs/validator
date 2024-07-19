// Package util contains utility functions for working with strings.
package util

import "strings"

// DeDupeStrSlice deduplicates a slices of strings
func DeDupeStrSlice(ss []string) []string {
	found := make(map[string]bool)
	l := []string{}
	for _, s := range ss {
		if _, ok := found[s]; !ok {
			found[s] = true
			l = append(l, s)
		}
	}
	return l
}

// Sanitize sends a string to lowercase, trims whitespace, replaces all non alphanumeric characters with a dash,
// removes consecutive dashes, and trims any leading or trailing dashes.
func Sanitize(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			return r
		}
		return '-'
	}, s)

	// Remove consecutive dashes
	var b strings.Builder
	prevDash := false
	for _, r := range s {
		if r == '-' {
			if !prevDash {
				b.WriteRune(r)
				prevDash = true
			}
		} else {
			b.WriteRune(r)
			prevDash = false
		}
	}

	return strings.Trim(b.String(), "-")
}
