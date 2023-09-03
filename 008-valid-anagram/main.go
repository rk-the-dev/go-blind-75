package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	// Convert the strings to lowercase to make it case-insensitive
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	// If the lengths of the two strings are different, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Convert the strings to slices of characters for comparison
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")

	// Sort the slices of characters
	sort.Strings(sChars)
	sort.Strings(tChars)

	// Compare the sorted slices of characters
	for i := 0; i < len(sChars); i++ {
		if sChars[i] != tChars[i] {
			return false
		}
	}

	return true
}
