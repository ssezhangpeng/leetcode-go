package leetcode_0073

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	rowHasZero, colHasZero := false, false // 第一行、列是否存在0

	for r := 0; r < m; r++ {
		if matrix[r][0] == 0 {
			colHasZero = true
			break
		}
	}
	for c := 0; c < n; c++ {
		if matrix[0][c] == 0 {
			rowHasZero = true
			break
		}
	}
	// 复用 matrix 的第一行和第一列
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}
	// 填充除第一行、列之外的 matrix
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}
	// 填充第一行、列
	if rowHasZero {
		for c := 0; c < n; c++ {
			matrix[0][c] = 0
		}
	}
	if colHasZero {
		for r := 0; r < m; r++ {
			matrix[r][0] = 0
		}
	}
}
