package leetcode_0115

// ref: 官方题解
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	// dp[i][j] 标识 s[i:] 中 t[j:] 的字串个数
	// dp[0][0] 即为所求
	dp := make([][]int, m+1)
	// 初始化
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	// 转移方程
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// 两种 case
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				// 一种 case
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
