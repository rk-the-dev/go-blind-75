package main

import "fmt"

func main() {
	s := "]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) > 0 {
				item := stack[len(stack)-1]
				if item != brackets[v] {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
