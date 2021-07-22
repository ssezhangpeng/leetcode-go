package leetcode_0084

func largestRectangleArea(heights []int) int {
	maxArea := 0
	var stk []int
	heights = append(heights, 0)

	for i := 0; i < len(heights); {
		if len(stk) == 0 {
			stk = append(stk, i)
			i++
		} else {
			topVal := stk[len(stk)-1]
			if heights[topVal] < heights[i] {
				stk = append(stk, i)
				i++
			} else {
				stk = stk[:len(stk)-1]
				width := i
				if len(stk) > 0 {
					l := stk[len(stk)-1]
					width = i - l - 1
				}
				maxArea = max(maxArea, heights[topVal]*width)
			}
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
