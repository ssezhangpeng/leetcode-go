package leetcode_0042

func trap(height []int) int {
	water := 0
	var stk []int

	for i := 0; i < len(height); {
		if len(stk) == 0 || height[i] <= height[stk[len(stk)-1]] {
			stk = append(stk, i)
			i++
		} else {
			down := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				continue
			}
			left := stk[len(stk)-1]
			wid := i - left - 1
			hei := min(height[left], height[i]) - height[down]
			water += wid * hei
		}
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
