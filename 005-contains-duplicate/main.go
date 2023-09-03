package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	traces := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := traces[v]; ok {
			return true
		} else {
			traces[v] = struct{}{}
		}
	}
	return false
}
