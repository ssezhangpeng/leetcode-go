package leetcode_0062

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// init
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] += dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
