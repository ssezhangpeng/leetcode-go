package leetcode_0044

// this algorithm is exceed time in leetcode
func isMatch(s string, p string) bool {
	return isMatchCore(s, p)
}

func isMatchCore(s, p string) bool {
	if len(p) > 0 && p[0] == '*' {
		for len(p) > 0 && p[0] == '*' {
			p = p[1:]
		}

		if len(p) == 0 {
			return true
		}

		for len(s) > 0 && !isMatchCore(s, p) {
			s = s[1:]
		}

		// if s == nil, p != nil, should return false
		return len(s) != 0
	} else if len(p) == 0 || len(s) == 0 {
		return len(p) == len(s)
	} else if p[0] == s[0] || p[0] == '?' {
		return isMatchCore(s[1:], p[1:])
	} else {
		return false
	}
}
