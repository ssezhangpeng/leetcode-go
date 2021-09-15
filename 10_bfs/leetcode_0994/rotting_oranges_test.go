package leetcode_0994

import "testing"

func TestRottingOranges(t *testing.T) {
	input := [][]int {{2, 1, 1},{1,1,0},{0,1,1}}

	ans := orangesRotting(input)

	if ans != 4 {
		t.Fatalf("want: %d, get: %d\n", 4, ans)
	}
}
