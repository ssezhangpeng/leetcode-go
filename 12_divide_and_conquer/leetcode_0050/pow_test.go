package leetcode_0050

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	var x float64 = 0.00001
	var n int = 2147483647

	ans := myPow(x, n)
	fmt.Println(ans)
}
