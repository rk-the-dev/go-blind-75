package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	result := twoSum(nums, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	trail := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		val, ok := trail[nums[i]]
		if ok {
			return []int{val, i}
		} else {
			trail[diff] = i
		}
	}
	return nil
}
