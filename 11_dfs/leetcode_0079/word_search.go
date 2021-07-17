package leetcode_0079

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if backtracking(board, i, j, 0, word, visited) {
					return true
				}
			}
		}
	}
	return false
}

func backtracking(board [][]byte, row, col, start int, word string, visited [][]bool) bool {
	m := len(board)
	n := len(board[0])

	if start == len(word) {
		return true
	}

	if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != word[start] || visited[row][col] {
		return false
	}

	visited[row][col] = true
	find := backtracking(board, row+1, col, start+1, word, visited) ||
		backtracking(board, row-1, col, start+1, word, visited) ||
		backtracking(board, row, col+1, start+1, word, visited) ||
		backtracking(board, row, col-1, start+1, word, visited)
	visited[row][col] = false

	return find
}
