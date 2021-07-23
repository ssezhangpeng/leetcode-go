package leetcode_0097

// ref: https://www.bilibili.com/video/BV1kK4y1575o?from=search&seid=13150942256442405875
func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	k := len(s3)

	if m+n != k {
		return false
	}

	// dp[i][j] 为 s[0->i] 和 s[0->j] 能交叉形成 s[0->i+j]
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if j == 0 {
				if s1[0:i] == s3[0:i] {
					dp[i][j] = true
				}
			} else if i == 0 {
				if s2[0:j] == s3[0:j] {
					dp[i][j] = true
				}
			} else {
				if s1[i-1] == s3[i+j-1] && dp[i-1][j] ||
					s2[j-1] == s3[i+j-1] && dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m][n]
}
