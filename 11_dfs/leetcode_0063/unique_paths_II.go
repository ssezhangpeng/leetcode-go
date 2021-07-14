package leetcode_0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	Stone := false
	for i := 0; i < m; i++ {
		if Stone || obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			Stone = true
		} else {
			dp[i][0] = 1
		}
	}
	Stone = false
	for j := 0; j < n; j++ {
		if Stone || obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
			Stone = true
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
