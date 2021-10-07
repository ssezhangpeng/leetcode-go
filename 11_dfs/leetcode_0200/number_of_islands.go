package leetcode_0200

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])

	// 递归出口
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
