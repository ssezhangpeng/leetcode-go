package leetcode_0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	len := len(nums)
	if len < 4 {
		return nil
	}
	ans := make([][]int, 0)

	sort.Ints(nums)

	for i:=0; i<len; i++ {
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		for j:=i+1; j<len; j++ {
			if j>i+1 && nums[j]==nums[j-1] {
				continue
			}
			l, r := j+1, len-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for ;l<r && nums[l]==nums[l-1];l++ {}
					for ;l<r && nums[r]==nums[r+1];r-- {}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ans
}
