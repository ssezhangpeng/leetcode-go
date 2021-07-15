package leetcode_0980

func uniquePathsIII(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	count := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, visited, &count)
			}
		}
	}
	return count
}

func dfs(grid [][]int, i, j int, visited [][]bool, count *int) {
	m := len(grid)
	n := len(grid[0])

	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == -1 {
		return
	}

	if grid[i][j] == 2 {
		if allVisited(grid, visited) {
			*count += 1
		}
		return
	}

	visited[i][j] = true
	dfs(grid, i+1, j, visited, count)
	dfs(grid, i-1, j, visited, count)
	dfs(grid, i, j+1, visited, count)
	dfs(grid, i, j-1, visited, count)
	visited[i][j] = false
}

func allVisited(grid [][]int, visited [][]bool) bool {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != -1 && grid[i][j] != 2 && !visited[i][j] {
				return false
			}
		}
	}
	return true
}
