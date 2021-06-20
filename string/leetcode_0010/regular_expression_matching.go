package leetcode_0010

func isMatch(s string, p string) bool {
	return isMatchCore(s, p)
}

func isMatchCore(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	// p != nil at this point, but maybe s == nil
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			return isMatchCore(s, p[2:]) || isMatchCore(s[1:], p[2:]) || isMatchCore(s[1:], p)
		} else {
			return isMatchCore(s, p[2:])
		}
	}

	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		return isMatchCore(s[1:], p[1:])
	}

	return false
}
