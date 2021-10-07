package leetcode_0542

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	queue := make([][]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				mat[i][j] = -1
			}
		}
	}

	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]

			if x >= 0 && x < rows && y >= 0 && y < cols && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	return mat
}
