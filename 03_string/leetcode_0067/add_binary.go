package leetcode_0067

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	carry := 0
	idx_a, idx_b := len(a)-1, len(b)-1

	ans := make([]string, max(len(a), len(b))+1)
	idx_k := len(ans) - 1

	for idx_a >= 0 || idx_b >= 0 {
		val_a, val_b := 0, 0
		if idx_a >= 0 {
			val_a, _ = strconv.Atoi(string(a[idx_a]))
		}
		if idx_b >= 0 {
			val_b, _ = strconv.Atoi(string(b[idx_b]))
		}

		val := (val_a + val_b + carry) % 2
		carry = (val_a + val_b + carry) / 2

		ans[idx_k] = fmt.Sprintf("%d", val)
		idx_k--

		idx_a--
		idx_b--
	}
	if carry > 0 {
		ans[0] = fmt.Sprintf("%d", carry)
	}
	return strings.Join(ans, "")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
