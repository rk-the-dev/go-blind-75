package main

func maxSubArray(nums []int) int {
	curr := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curr = max(nums[i], nums[i]+curr)
		maxSum = max(maxSum, curr)
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
