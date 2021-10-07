package leetcode_0784

func letterCasePermutation(s string) []string {
	path := []byte(s)

	var ans []string

	backTracking(path, 0, &ans)

	return ans
}

func backTracking(path []byte, start int, ans *[]string) {
	// 递归出口
	*ans = append(*ans, string(path))

	for i := start; i < len(path); i++ {
		if isDigit(path[i]) {
			continue
		}

		if isUpper(path[i]) {
			path[i] += 32
		} else {
			path[i] -= 32
		}

		backTracking(path, i+1, ans)

		if isUpper(path[i]) {
			path[i] += 32
		} else {
			path[i] -= 32
		}
	}
}

func isDigit(ch byte) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}

func isUpper(ch byte) bool {
	if 'A' <= ch && ch <= 'Z' {
		return true
	}
	return false
}
