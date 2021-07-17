package leetcode_0037

func solveSudoku(board [][]byte) {
	backtracking(board)
	return
}

func backtracking(board [][]byte) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isValid(board, i, j, byte(k)) {
					board[i][j] = byte(k)
					if backtracking(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, val byte) bool {
	m := len(board)
	n := len(board[0])

	// 判断行重复
	for j := 0; j < n; j++ {
		if board[row][j] == val {
			return false
		}
	}
	// 判断列重复
	for i := 0; i < m; i++ {
		if board[i][col] == val {
			return false
		}
	}
	// 判断小矩形重复
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}
