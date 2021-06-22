package leetcode_0038

import "strconv"

func countAndSay(n int) string {
	s := "1"
	for n-1 > 0 {
		s = getNext(s)
		n--
	}
	return s
}

func getNext(s string) string {
	ans := ""
	times := 0
	l, r := 0, 0
	for ; r < len(s); r++ {
		if s[r] != s[l] {
			times = r - l
			ans = ans + strconv.Itoa(times) + string(s[l])
			l = r
		}
	}
	if r == len(s) {
		times = r - l
		ans = ans + strconv.Itoa(times) + string(s[l])
	}
	return ans
}
