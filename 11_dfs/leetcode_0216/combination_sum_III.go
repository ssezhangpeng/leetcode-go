package leetcode_0216

func combinationSum3(k int, n int) [][]int {
	candidates := make([]int, 9)
	for i := 1; i <= 9; i++ {
		candidates[i-1] = i
	}

	var ans [][]int
	var path []int

	backtracking(candidates, 0, 0, n, k, path, &ans)

	return ans
}

func backtracking(candidates []int, start, sum, target, k int, path []int, ans *[][]int) {
	if sum == target && len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	// 剪枝
	if sum > target || len(path) > k {
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtracking(candidates, i+1, sum+candidates[i], target, k, path, ans)
		path = path[:len(path)-1]
	}
}
