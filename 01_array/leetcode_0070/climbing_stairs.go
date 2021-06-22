package leetcode_0070

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	// 1, 1, 2, 3, 5
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}
	return curr
}
