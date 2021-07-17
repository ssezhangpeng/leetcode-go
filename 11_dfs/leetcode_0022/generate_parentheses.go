package leetcode_0022

import "strings"

func generateParenthesis(n int) []string {
	var ans []string
	var path []string

	backtracking(n, n, 0, n, path, &ans)

	return ans
}

func backtracking(left, right, start, n int, path []string, ans *[]string) {
	if left == 0 && right == 0 {
		s := strings.Join(path, "")
		*ans = append(*ans, s)
		return
	}

	if left < 0 || right < 0 || left > right {
		return
	}

	for i := start; i < 2*n; i++ {
		path = append(path, "(")
		backtracking(left-1, right, i+1, n, path, ans)
		path = path[:len(path)-1]

		path = append(path, ")")
		backtracking(left, right-1, i+1, n, path, ans)
		path = path[:len(path)-1]
	}
}
