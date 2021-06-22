package leetcode_0037

// ref: https://www.cnblogs.com/grandyang/p/4421852.html
func solveSudoku(board [][]byte) {
	solveSudokuCore(board, 0, 0)
}

func solveSudokuCore(board [][]byte, i, j int) bool {
	//递归出口
	if i == 9 {
		return true
	}

	if j >= 9 {
		return solveSudokuCore(board, i+1, 0)
	}
	if board[i][j] != '.' {
		return solveSudokuCore(board, i, j+1)
	}
	for c := '1'; c <= '9'; c++ {
		if !isValid(board, i, j, byte(c)) {
			continue
		}
		board[i][j] = byte(c)
		if solveSudokuCore(board, i, j+1) {
			return true
		}
		// backTrack
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, c byte) bool {
	// check row
	for col := 0; col < 9; col++ {
		if board[i][col] == c {
			return false
		}
	}
	// check col
	for row := 0; row < 9; row++ {
		if board[row][j] == c {
			return false
		}
	}
	// check box
	tr, tc := i-i%3, j-j%3
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row+tr][col+tc] == c {
				return false
			}
		}
	}
	return true
}
