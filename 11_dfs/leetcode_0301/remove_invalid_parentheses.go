package leetcode_0301

//TODO: 未完成
func removeInvalidParentheses(s string) []string {
	length := len(s)
	l, r := 0, 0

	for i := 0; i < length; i++ {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' && l == 0 {
			r++
		} else {
			l--
		}
	}

	var ans []string
	dfs(s, 0, l, r, &ans)

	return ans
}

func dfs(s string, start, l, r int, ans *[]string) {
	// 递归出口
	if l == 0 && r == 0 {
		if isValid(s) {
			*ans = append(*ans, s)
			return
		}
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}

		ch := s[i]
		curr := remove(s, i)

		if ch == ')' && r > 0 {
			dfs(curr, i, l, r-1, ans)
		}
		if ch == '(' && l > 0 {
			dfs(curr, i, l-1, r, ans)
		}
	}
}

func isValid(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return true
}

func remove(s string, index int) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if i != index {
			ans += string(s[i])
		}
	}
	return ans
}
