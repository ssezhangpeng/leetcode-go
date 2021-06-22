package leetcode_0065

// ref: https://leetcode-cn.com/problems/valid-number/solution/you-xiao-shu-zi-by-leetcode-solution-298l/#comment

func isNumber(s string) bool {
	numFlag, dotFlag, eFlag := false, false, false

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false // reJudge integer after 'e' or 'E'
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	// avoid case: s == '.' or 'e/E' or '+/-' etc...
	// string s must have num
	return numFlag
}
