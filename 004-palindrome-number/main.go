package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	int_string := strconv.Itoa(x)
	if len(int_string) <= 1 {
		return true
	}
	for i, j := 0, len(int_string)-1; i <= j; i, j = i+1, j-1 {
		if int_string[i] != int_string[j] {
			return false
		}
	}
	return true
}
func main() {
	num := 1
	fmt.Println(isPalindrome(num))
}
