package leetcode_0036

func isValidSudoku(board [][]byte) bool {
	rowBuf := make([][9]bool, 9) // rowBuf[r][val] 表示 val 在 r 行中是否出现过
	colBuf := make([][9]bool, 9) // colBuf[c][val] 表示 val 在 c 列中是否出现过
	boxBuf := make([][9]bool, 9) // boxBuf[i][val] 表示 val 在 i 个小格子中是否出现过

	// traver && add buff
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				val := board[r][c] - '1'
				if rowBuf[r][val] || colBuf[c][val] || boxBuf[r/3*3+c/3][val] {
					return false
				}
				rowBuf[r][val] = true
				colBuf[c][val] = true
				boxBuf[r/3*3+c/3][val] = true
			}
		}
	}
	return true
}
