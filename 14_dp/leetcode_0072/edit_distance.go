package leetcode_0072

/*
	/* dp[i][j] 标识 A[0, i] 和 B[0, j] 之间的最小编辑距离
	/* 如果 c == d, 则 dp[i][j] = dp[i-1][j-1]
	/* 如果 c != d, 则
	/* a) 如果将 c 替换为 d,   则 dp[i][j] = dp[i][j-1] + 1
	/* b) 如果将 c 后面添加 d, 则 dp[i][j] = dp[i][j-1] + 1
	/* c) 如果将 c 删除,      则 dp[i][j] = dp[i-1][j] + 1
*/

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	// dp[i][j] 标识 word1[0->i]到 word2[0->j] 的最小编辑距离
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	// 动态转移方程
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				minChange := min(dp[i-1][j], dp[i][j-1])
				dp[i][j] = 1 + min(dp[i-1][j-1], minChange)
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
