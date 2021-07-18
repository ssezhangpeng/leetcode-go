package leetcode_0003

func lengthOfLongestSubstring(s string) int {
	visited := make([]int, 256)

	for i := 0; i < 256; i++ {
		visited[i] = -1
	}

	ans := 0
	l := 0 // window: [l, r]
	for r := 0; r < len(s); r++ {
		if visited[s[r]] >= l { // 在 window[l, r] 内出现重复元素
			ans = max(ans, r-l)
			l = visited[s[r]] + 1
		}
		visited[s[r]] = r
	}

	// 处理最后一次的情况
	return max(ans, len(s)-l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
