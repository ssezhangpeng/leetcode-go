package leetcode_0994

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// 超级源点方式
	fresh := 0
	queue := make([][]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh += 1
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	// 开始 BFS 模板
	level := 0
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]

			for _, v := range direction {
				x := point[0] + v[0]
				y := point[1] + v[1]

				if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != 1 {
					continue
				}

				grid[x][y] = 2
				queue = append(queue, []int{x, y})

				fresh -= 1
			}
		}
		// 每跑完一圈，距离加一，如果没有未腐烂的水果，则直接返回
		level += 1
		if fresh == 0 {
			break
		}
	}

	if fresh > 0 {
		return -1
	}

	return level
}
