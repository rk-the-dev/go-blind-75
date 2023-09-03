package main

func removeDuplicates(nums []int) int {
	traces := map[int]struct{}{}
	unique := []int{}
	for _, v := range nums {
		if _, ok := traces[v]; !ok {
			unique = append(unique, v)
		} else {
			traces[v] = struct{}{}
		}
	}
	return len(unique)
}
