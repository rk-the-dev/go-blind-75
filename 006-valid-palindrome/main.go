package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	strArr := []rune(clearedStr)
	for i, j := 0, len(strArr)-1; i <= j; i, j = i+1, j-1 {
		if strArr[i] != strArr[j] {
			return false
		}
	}
	return true
}

func main() {
	s := ""
	result := isPalindrome(s)
	fmt.Println(result)
}
