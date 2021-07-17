package leetcode_0036

func isValidSudoku(board [][]byte) bool {
	rowBuf := make([][9]bool, 9)
	colBuf := make([][9]bool, 9)
	boxBuf := make([][9]bool, 9)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}

			val := board[row][col] - '1'
			if rowBuf[row][val] || colBuf[col][val] || boxBuf[row/3*3+col/3][val] {
				return false
			}

			rowBuf[row][val] = true
			colBuf[col][val] = true
			boxBuf[row/3*3+col/3][val] = true
		}
	}
	return true
}
