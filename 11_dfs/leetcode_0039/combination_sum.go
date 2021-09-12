package leetcode_0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int

	sort.Ints(candidates)
	backtracking(candidates, 0, target, 0, path, &ans)

	return ans
}

func backtracking(candidates []int, sum, target, start int, path []int, ans *[][]int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtracking(candidates, sum+candidates[i], target, i, path, ans)
		path = path[:len(path)-1]
	}
}
