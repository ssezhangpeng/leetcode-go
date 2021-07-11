package leetcode_0047

import "sort"

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var path []int
	visited := make([]bool, len(nums))

	sort.Ints(nums)
	backtracking(nums, path, &ans, visited)

	return ans
}

func backtracking(nums, path []int, ans *[][]int, visited []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// de-duplicate
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}

		path = append(path, nums[i])
		visited[i] = true

		backtracking(nums, path, ans, visited)

		path = path[:len(path)-1]
		visited[i] = false
	}
}
