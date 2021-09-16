package leetcode_0198

func rob(nums []int) int {
	len := len(nums)
	dp := make([]int, len)

	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return max(nums[0], nums[1])
	}

	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i:=2; i<len; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}

	return dp[len-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
