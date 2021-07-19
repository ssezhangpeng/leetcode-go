package leetcode_0053

func maxSubArray(nums []int) int {
	dp := nums
	maxVal := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxVal = max(maxVal, dp[i])
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
