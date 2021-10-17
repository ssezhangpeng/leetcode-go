package leetcode_0330

func lengthOfLIS(nums []int) int {
	size := len(nums)

	// dp[i] 标识以下标 i 开头的最长子序列
	dp := make([]int, size)

	// init
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	maxLength := 1
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				maxLength = max(maxLength, dp[i])
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
