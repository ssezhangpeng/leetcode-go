package leetcode_0130

func solve_dfs(board [][]byte) {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(&board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(&board, i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(&board, 0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(&board, m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

	return
}

func dfs(board *[][]byte, i, j int) {
	m := len(*board)
	n := len((*board)[0])

	if i < 0 || i >= m || j < 0 || j >= n || (*board)[i][j] != 'O' {
		return
	}

	(*board)[i][j] = 'A'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}
