package leetcode0438

func findAnagrams(s string, p string) []int {
	lens := len(s)
	lenp := len(p)

	vs := make([]int, 26)
	vp := make([]int, 26)

	l, r := 0, 0

	// 初始化
	for _, ch := range p {
		vp[ch-'a']++
	}

	var result []int
	// 窗口开始滑动
	for r < lens {
		if r >= lenp {
			vs[s[l]-'a']--
			l++
		}

		vs[s[r]-'a']++
		r++

		if isEqual(vs, vp) {
			result = append(result, l)
		}
	}

	return result
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
