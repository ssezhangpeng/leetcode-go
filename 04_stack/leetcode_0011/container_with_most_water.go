package leetcode_0011

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1

	for l < r {
		w := r - l
		h := min(height[l], height[r])
		maxArea = max(maxArea, w*h)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
