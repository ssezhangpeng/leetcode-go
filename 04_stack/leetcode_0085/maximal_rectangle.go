package leetcode_0085

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	ans := 0
	m, n := len(matrix), len(matrix[0])

	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		ans = max(ans, largestRectangleArea(height))
	}

	return ans
}

// leetcode_0084
func largestRectangleArea(heights []int) int {
	maxArea := 0
	var incrStack []int

	// 最后追加一个 0, 是为了处理递增的 heights, 这样最后也能进行一次收尾处理
	heights = append(heights, 0)
	for i:=0; i<len(heights); {
		if len(incrStack) == 0 || heights[incrStack[len(incrStack)-1]] < heights[i] {
			incrStack = append(incrStack, i)
			i++
		} else {
			index := incrStack[len(incrStack)-1]
			incrStack = incrStack[:len(incrStack)-1]

			currHeight := heights[index]
			currWidth := i
			if len(incrStack) > 0 {
				currWidth = i - incrStack[len(incrStack)-1] - 1
			}

			maxArea = max(maxArea, currWidth * currHeight)
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
