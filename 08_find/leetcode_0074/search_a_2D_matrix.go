package leetcode_0074

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	row, col := 0, n-1
	for {
		if row >= m || col < 0 {
			return false
		}
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
}
