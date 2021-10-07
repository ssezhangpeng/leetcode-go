package leetcode_0056

import "sort"

func merge(intervals [][]int) [][]int {
	var ans [][]int

	// sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		current := intervals[i]
		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			ans = append(ans, current)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
