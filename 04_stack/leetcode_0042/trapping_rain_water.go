package leetcode_0042

func trap(height []int) int {
	var water int
	var stack []int

	for i := 0; i < len(height); i++ {
		// 尝试持续出栈
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				continue
			}

			left := stack[len(stack)-1]
			h := min(height[i], height[left]) - height[top]
			w := i - left - 1

			water += h * w
		}
		stack = append(stack, i)
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
