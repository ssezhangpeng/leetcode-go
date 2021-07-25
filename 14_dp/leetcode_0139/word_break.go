package leetcode_0139

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	m := make(map[string]bool)

	for _, word := range wordDict {
		m[word] = true
	}

	dp := make([]bool, size+1)
	dp[0] = true

	for j := 1; j <= size; j++ {
		for i := 0; i < j; i++ {
			ss := s[i:j]
			if dp[i] && m[ss] {
				dp[j] = true
				break
			}
		}
	}
	return dp[size]
}
