package leetcode0347

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	arr := []int{4, 1, -1, 2, -1, 2, 3}
	k := 2

	ans := topKFrequent(arr, k)
	fmt.Printf("ans: %+v\n", ans)
}
