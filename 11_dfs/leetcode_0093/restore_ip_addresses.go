package leetcode_0093

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var ans []string
	var path []string

	backtracking(s, 0, path, &ans)

	return ans
}

func backtracking(s string, start int, path []string, ans *[]string) {
	if len(path) == 4 && start == len(s) {
		s := strings.Join(path, ".")
		*ans = append(*ans, s)
		return
	}

	for i := start; i < start+3; i++ {
		if i >= len(s) {
			return
		}
		// 剪枝
		if !isValid(s[start : i+1]) {
			return
		}

		path = append(path, s[start:i+1])
		backtracking(s, i+1, path, ans)
		path = path[:len(path)-1]
	}
}

func isValid(s string) bool {
	// 判断前导零
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	num, _ := strconv.Atoi(s)
	return num <= 255
}
