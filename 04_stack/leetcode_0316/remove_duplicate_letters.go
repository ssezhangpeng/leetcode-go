package leetcode_0316

func removeDuplicateLetters(s string) string {
	size := len(s)
	var stk []byte

	visited := make(map[byte]struct{})
	last := make([]int, 128)
	for i := 0; i < size; i++ {
		last[s[i]] = i
	}

	for i := 0; i < size; i++ {
		c := s[i]
		if _, ok := visited[c]; ok {
			continue
		}

		for len(stk) > 0 && c <= stk[len(stk)-1] && i <= last[stk[len(stk)-1]] {
			delete(visited, stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}

		stk = append(stk, c)
		visited[c] = struct{}{}
	}

	ans := ""
	for i := 0; i < len(stk); i++ {
		ans += string(stk[i])
	}

	return ans
}
