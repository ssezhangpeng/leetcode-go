package leetcode_0005

func longestPalindrome(s string) string {
	length := len(s)
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	ans := ""
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i <= 2 || dp[i+1][j-1])
			if dp[i][j] && (j-i+1 > len(ans)) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
