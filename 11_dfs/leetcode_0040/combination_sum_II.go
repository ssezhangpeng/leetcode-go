package leetcode_0040

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int

	sort.Ints(candidates)
	backtracking(candidates, 0, 0, target, path, &ans)

	return ans
}

func backtracking(candidates []int, start, sum, target int, path []int, ans *[][]int) {
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
		// de-duplicate
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtracking(candidates, i+1, sum+candidates[i], target, path, ans)
		path = path[:len(path)-1]
	}
}
