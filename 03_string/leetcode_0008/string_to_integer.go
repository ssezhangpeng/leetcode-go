package leetcode_0008

import (
	"math"
)

func myAtoi(s string) int {
	idx := 0
	flag := int64(1)
	ans := int64(0)

	for ; idx < len(s); idx++ {
		if s[idx] != ' ' {
			break
		}
	}
	if idx == len(s) {
		return int(ans)
	}

	if s[idx] == '+' || s[idx] == '-' {
		if s[idx] == '-' {
			flag = -1
		}
		idx++
	}
	if idx == len(s) {
		return int(ans)
	}

	for ; idx < len(s); idx++ {
		if '0' <= s[idx] && s[idx] <= '9' {
			ans = ans*10 + int64(s[idx]-'0')
			if flag*ans < math.MinInt32 {
				return math.MinInt32
			}
			if flag*ans > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	return int(flag * ans)
}
