package leetcode_0130

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

type Pos struct {
	X int
	Y int
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	queue := make([]Pos, 0)
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, Pos{i, 0})
		}
		if board[i][n-1] == 'O' {
			queue = append(queue, Pos{i, n - 1})
		}
	}

	for j := 1; j < n-1; j++ {
		if board[0][j] == 'O' {
			queue = append(queue, Pos{0, j})
		}
		if board[m-1][j] == 'O' {
			queue = append(queue, Pos{m - 1, j})
		}
	}

	for len(queue) > 0 {
		x, y := queue[0].X, queue[0].Y
		queue = queue[1:]

		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || mx >= m || my < 0 || my >= n || board[mx][my] != 'O' {
				continue
			}
			queue = append(queue, Pos{mx, my})
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
