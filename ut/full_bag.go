package ut

import "fmt"

func fullBag(weight []int, value []int, bag int) int {
	// dp[i][j]: 表示从[0, i]的物品中重复选取，放入容量为j的背包中，获取的最大价值
	dp := make([][]int, len(weight)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bag+1)
	}

	// // init dp
	// for i := 0; i < len(weight); i++ {
	// 	dp[i][0] = 0
	// }
	// for j := weight[0]; j <= bag; j++ {
	// 	dp[0][j] = value[0]
	// }

	// 动态转移方程
	for i := 1; i <= len(weight); i++ {
		for j := weight[i-1]; j <= bag; j++ {
			// NOTE: 这里和 0-1 背包不同, 当第i个物品装的时候，还是从[0, i]的物品中选取，放入容量为 j-weight[i] 的背包中
			// 在 0-1 背包中，还是从[0, i-1]的物品中选取，放入容量为 j-weight[i] 的背包中
			dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i-1]]+value[i-1])
		}
	}

	fmt.Printf("%v\n", dp)

	return dp[len(weight)][bag]
}
