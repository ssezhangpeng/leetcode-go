package leetcode_0052

import "strings"

func totalNQueens(n int) int {
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = strings.Repeat(".", n)
	}

	count := 0
	backtracking(grid, 0, &count)

	return count
}

func backtracking(grid []string, row int, count *int) {
	if row == len(grid) {
		*count += 1
		return
	}

	for col := 0; col < len(grid); col++ {
		if isValid(grid, row, col) {
			s := []byte(grid[row])
			s[col] = 'Q'
			grid[row] = string(s)

			backtracking(grid, row+1, count)

			s[col] = '.'
			grid[row] = string(s)
		}
	}
}

func isValid(grid []string, row, col int) bool {
	m, n := len(grid), len(grid)

	// 判断相同列
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
