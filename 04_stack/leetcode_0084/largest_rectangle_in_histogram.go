package leetcode_0084

import "github.com/emirpasic/gods/stacks/linkedliststack"

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stk := linkedliststack.New()
	heights = append(heights, 0)

	for i := 0; i < len(heights); {
		if stk.Empty() {
			stk.Push(i)
			i++
		} else {
			topVal, _ := stk.Peek()
			if heights[topVal.(int)] < heights[i] {
				stk.Push(i)
				i++
			} else {
				stk.Pop()
				width := i
				if !stk.Empty() {
					l, _ := stk.Peek()
					width = i - l.(int) - 1
				}
				maxArea = max(maxArea, heights[topVal.(int)]*width)
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
