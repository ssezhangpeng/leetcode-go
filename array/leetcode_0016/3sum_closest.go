package leetcode_0016

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return 0 - x
}

func threeSumClosest(nums []int, target int) int {
	len := len(nums)
	if len < 3 {
		return 0
	}

	minDist := math.MaxInt8
	ans := 0
	sort.Ints(nums)
	for i:=0; i<len; i++ {
		l, r := i+1, len-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			dist := abs(sum - target)
			if dist < minDist {
				minDist = dist
				ans = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
