package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	result := maxProfit(prices)
	fmt.Println(result)
}
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
