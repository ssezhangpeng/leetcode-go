package leetcode_0048

func rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	tr, tc, dr, dc := 0, 0, m-1, n-1
	for tr < dr {
		rotateCore(matrix, tr, tc, dr, dc)
		tr++
		tc++
		dr--
		dc--
	}
}

func rotateCore(matrix [][]int, tr, tc, dr, dc int) {
	times := dc - tc
	for i := 0; i < times; i++ {
		temp := matrix[tr][tc+i]
		matrix[tr][tc+i] = matrix[dc-i][tc]
		matrix[dc-i][tc] = matrix[dr][dc-i]
		matrix[dr][dc-i] = matrix[tc+i][dc]
		matrix[tc+i][dc] = temp
	}
}
