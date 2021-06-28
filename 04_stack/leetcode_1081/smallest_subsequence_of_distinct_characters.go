package leetcode_1081

func smallestSubsequence(s string) string {
	size := len(s)
	var stk []byte

	visited := make([]bool, 128)
	last := make([]int, 128)

	for i := 0; i < size; i++ {
		last[s[i]] = i
	}

	for i := 0; i < size; i++ {
		c := s[i]

		if visited[c] {
			continue
		}

		for len(stk) > 0 && c <= stk[len(stk)-1] && i <= last[stk[len(stk)-1]] {
			visited[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}

		stk = append(stk, c)
		visited[c] = true
	}

	ans := ""
	for i := 0; i < len(stk); i++ {
		ans += string(stk[i])
	}
	return ans
}
