package leetcode_0070

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	before, after := 1, 2
	for i:=3; i<=n; i++ {
		before, after = after, before + after
	}

	return after
}
