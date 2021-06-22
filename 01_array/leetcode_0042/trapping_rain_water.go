package leetcode_0042

import "math"

func trap(height []int) int {
	len := len(height)

	lMax := math.MinInt8
	left := make([]int, len)
	for i := 1; i < len; i++ {
		lMax = max(lMax, height[i-1])
		left[i] = lMax
	}
	rMax := math.MinInt8
	right := make([]int, len)
	for i := len - 2; i >= 0; i-- {
		rMax = max(rMax, height[i+1])
		right[i] = rMax
	}

	water := 0
	for i := 1; i < len-1; i++ {
		minHeight := min(left[i], right[i])
		if minHeight > height[i] {
			water += minHeight - height[i]
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
