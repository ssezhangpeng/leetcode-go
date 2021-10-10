package leetcode_0240

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	i, j := 0, n-1
	for true {
		if i >= m || j < 0 {
			return false
		}

		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
