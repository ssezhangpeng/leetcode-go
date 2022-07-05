package ut

func zeroOneBag(weight []int, value []int, bag int) int {
	// dp[i][j]: 表示从[0, i]中挑选物品，放入容量为 j 的背包中，多能获取的最大价值
	dp := make([][]int, len(weight))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bag+1)
	}

	// init dp
	for i := 0; i < len(weight); i++ {
		dp[i][0] = 0
	}
	for j := weight[0]; j <= bag; j++ {
		dp[0][j] = value[0]
	}

	// 动态转移方程
	for i := 1; i < len(weight); i++ {
		for j := weight[i]; j <= bag; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	// fmt.Printf("dp: %v\n", dp)

	return dp[len(weight)-1][bag]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
