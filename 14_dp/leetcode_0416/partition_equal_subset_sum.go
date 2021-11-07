package leetcode0416

func canPartition(nums []int) bool {
	// 0-1 背包问题
	// https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/0-1-bei-bao-wen-ti-bian-ti-zhi-zi-ji-fen-ge-by-lab/
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	n := len(nums)
	target := sum / 2

	dp := make([][]bool, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
	}

	// 初始化
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	// 状态转移方程
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i-1] < 0 { //note: 第 i 个物品的质量应该是 nums[i-1]
				// 背包容量不足, 不放入
				dp[i][j] = dp[i-1][j]
			} else {
				// 背包容量充足，可不装入可装入
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][target]
}
