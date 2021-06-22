package leetcode_0015

import "sort"

func threeSum(nums []int) [][]int {
	len := len(nums)
	if len < 3 {
		return nil
	}

	sort.Ints(nums)

	ans := make([][]int, 0)
	for i := 0; i < len; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
