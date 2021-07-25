package leetcode_0091

func numDecodings(s string) int {
	n := len(s)
	// dp[i] 标识前 i 个字符的解码方式
	// dp[n] 即为所求
	dp := make([]int, n+1)

	// 初始化
	dp[0] = 1

	for i := 1; i <= n; i++ {
		// case1: 只用 dp[i] 一位进行编码
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// case2: 用 dp[i-1], dp[i] 两位进行编码
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
