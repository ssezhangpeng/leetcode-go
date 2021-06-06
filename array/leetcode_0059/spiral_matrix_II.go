package leetcode_0059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	num := 1
	tr, tc, dr, dc := 0, 0, n-1, n-1
	for tr <= dr {
		generateMatrixCore(matrix, tr, tc, dr, dc, &num)
		tr++
		tc++
		dr--
		dc--
	}
	// handler single when n is odd
	if n%2 != 0 {
		matrix[tr][tc] = num
	}
	return matrix
}

func generateMatrixCore(matrix [][]int, tr, tc, dr, dc int, num *int) {
	// left --> right
	for j := tc; j < dc; j++ {
		matrix[tr][j] = *num
		*num++
	}
	// up --> down
	for i := tr; i < dr; i++ {
		matrix[i][dc] = *num
		*num++
	}
	// right --> left
	for j := dc; j > tc; j-- {
		matrix[dr][j] = *num
		*num++
	}
	// down --> up
	for i := dr; i > tr; i++ {
		matrix[i][tc] = *num
		*num++
	}
}
