package leetcode_0058

import "strings"

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")

	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) == 0 {
			continue
		} else {
			return len(arr[i])
		}
	}
	return 0
}
