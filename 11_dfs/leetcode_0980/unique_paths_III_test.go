package leetcode_0980

import (
	"fmt"
	"testing"
)

func TestUniquePathsIII(t *testing.T) {
	grid := [][]int{
		{0, 1},
		{2, 0},
	}

	count := uniquePathsIII(grid)
	fmt.Println("count: ", count)
}
