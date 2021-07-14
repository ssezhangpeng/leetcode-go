package leetcode_0062

// dfs: exceed timeout when m=51 && n == 9
func uniquePathsDFS(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	return uniquePathsDFS(m-1, n) + uniquePathsDFS(m, n-1)
}

func uniquePaths(m int, n int) int {
	// dp[i][j] 为 matrix[0][0]-->matrix[i][j]的路径数目
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
