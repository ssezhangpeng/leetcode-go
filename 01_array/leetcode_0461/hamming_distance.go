package leetcode0461

func hammingDistance(x int, y int) int {
	return OneOfNum(x ^ y)
}

func OneOfNum(n int) int {
	cnt := 0

	for n != 0 {
		cnt++
		n = n & (n - 1)
	}

	return cnt
}
