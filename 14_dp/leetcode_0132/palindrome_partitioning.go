package leetcode_0132

func minCut(s string) int {
	palindromicGrid := PalindromicGrid(s)

	// dp[i] 标识 s[0:i] 的最少切割次数
	m := len(s)
	dp := make([]int, m)

	for i := 0; i < m; i++ {
		dp[i] = i
	}

	for i := 1; i < len(s); i++ {
		if palindromicGrid[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if palindromicGrid[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[m-1]
}

func PalindromicGrid(s string) [][]bool {
	m, n := len(s), len(s)
	dp := make([][]bool, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := i; j < m; j++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
