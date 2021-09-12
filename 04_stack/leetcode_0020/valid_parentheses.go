package leetcode_0020

func isValid(s string) bool {
	stk := make([]byte, 0)

	for i:=0; i<len(s); i++ {
		if len(stk) == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk = append(stk, s[i])
		} else {
			ch := stk[len(stk)-1]
			if (s[i] == ')' && ch == '(') || (s[i] == ']' && ch == '[') || (s[i] == '}' && ch == '{') {
				stk = stk[:len(stk)-1]
			} else {
				return false
			}
		}
	}

	return len(stk) == 0
}
