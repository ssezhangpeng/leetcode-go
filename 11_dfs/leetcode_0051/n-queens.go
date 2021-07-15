package leetcode_0051

import "strings"

func solveNQueens(n int) [][]string {
	var ans [][]string

	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = strings.Repeat(".", n)
	}

	backtracking(0, grid, &ans)

	return ans
}

func backtracking(row int, grid []string, ans *[][]string) {
	if row == len(grid) {
		temp := make([]string, len(grid))
		copy(temp, grid)
		*ans = append(*ans, temp)
		return
	}

	for col := 0; col < len(grid); col++ {
		if isValid(grid, row, col) {
			s := []byte(grid[row])
			s[col] = 'Q'
			grid[row] = string(s)

			backtracking(row+1, grid, ans)

			s[col] = '.'
			grid[row] = string(s)
		}
	}
}

func isValid(grid []string, row, col int) bool {
	m, n := len(grid), len(grid)

	// 判断当前列
	for i := 0; i < m; i++ {
		if grid[i][col] == 'Q' {
			return false
		}
	}

	// 判断主对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if grid[i][j] == 'Q' {
			return false
		}
	}

	// 判断副对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if grid[i][j] == 'Q' {
			return false
		}
	}
	return true
}
