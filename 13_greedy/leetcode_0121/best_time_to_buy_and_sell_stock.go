package leetcode_0121

import "math"

func maxProfit(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt32

	for _, price := range prices {
		maxProfit = max(maxProfit, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
