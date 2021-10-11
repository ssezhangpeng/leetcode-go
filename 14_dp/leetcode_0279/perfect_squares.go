package leetcode_0279

func numSquares(n int) int {
	// dp[i] 标识和为 i 所需要的最少完全平方数个数
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i*i <= n; i++ {
		// i*i 自己就是完全平方数
		dp[i*i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
