package leetcode_0739

func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	var stk []int

	ans := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		for len(stk) > 0 && temperatures[i] >= temperatures[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stk[len(stk)-1] - i
		}

		stk = append(stk, i)
	}

	return ans
}
