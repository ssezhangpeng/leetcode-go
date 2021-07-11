package leetcode_0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var path []int

	// do sort before de-duplicate
	sort.Ints(nums)
	backtracking(nums, 0, path, &ans)

	return ans
}

func backtracking(nums []int, startIndex int, path []int, ans *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ans = append(*ans, temp)

	for i := startIndex; i < len(nums); i++ {
		// de-duplicate
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		backtracking(nums, i+1, path, ans)

		path = path[:len(path)-1]
	}
}
