package leetcode_0122

func maxProfit(prices []int) int {
	sum := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			sum += profit
		}
	}

	return sum
}
