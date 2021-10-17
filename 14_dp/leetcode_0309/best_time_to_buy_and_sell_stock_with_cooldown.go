package leetcode_0309

func maxProfit(prices []int) int {
	size := len(prices)

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, 3)
	}

	// f[i][0]:第 i 天结束的时候手上有股票的最大收益
	// f[i][1]:第 i 天结束的时候手上无股票，并且处于冷冻期的最大收益
	// f[i][2]:第 i 天结束的时候手上无股票，并且不在冷冻期的最大收益

	// init
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	// 转移方程
	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[size-1][1], dp[size-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
