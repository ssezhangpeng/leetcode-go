package leetcode_0120

func minimumTotal(triangle [][]int) int {
	m := len(triangle)

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
