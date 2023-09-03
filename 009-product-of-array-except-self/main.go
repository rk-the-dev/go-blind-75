package main

func productExceptSelf(nums []int) []int {
	preficProduct := make([]int, len(nums))
	preficProduct[0] = 1
	for i := 1; i <= len(nums)-1; i++ {
		preficProduct[i] = preficProduct[i-1] * nums[i-1]
	}
	sufixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		preficProduct[i] = preficProduct[i] * sufixProduct
		sufixProduct = sufixProduct * nums[i]
	}
	return preficProduct
}
