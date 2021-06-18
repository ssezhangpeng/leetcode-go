package leetcode_0128

import "math"

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	used := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
		used[nums[i]] = false
	}

	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		currLength := 1

		for j := nums[i] + 1; m[j] && !used[j]; j++ {
			used[j] = true
			currLength++
		}

		for j := nums[i] - 1; m[j] && !used[j]; j-- {
			used[j] = true
			currLength++
		}

		maxLength = int(math.Max(float64(maxLength), float64(currLength)))
	}

	return maxLength
}
