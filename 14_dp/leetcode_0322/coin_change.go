package leetcode0322

func coinChange(coins []int, amount int) int {
	maxVal := amount + 1
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = maxVal
	}

	// init
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == maxVal {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
