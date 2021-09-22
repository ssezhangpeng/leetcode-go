package leetcode_0076

import "math"

func minWindow(s string, t string) string {
	window := make(map[uint8]int)
	need := make(map[uint8]int)

	// 记录匹配字符串中的元素和对应个数
	for i:=0; i<len(t); i++ {
		need[t[i]]++
	}

	// 滑动窗口进行匹配
	left, right := 0, 0
	// 记录最终结果
	start, minLength := 0, math.MaxInt16
	// 记录满足匹配的元素个数
	valid := 0

	for right < len(s) {
		ch := s[right]
		right++

		if _, ok := need[ch]; ok {
			window[ch]++
			if window[ch] == need[ch] {
				valid++
			}
		}

		// 如果找到一个覆盖区间，则尝试寻找下一个更短的区间
		for valid == len(need) {
			if right - left < minLength {
				start = left
				minLength = right - left
			}

			ch = s[left]
			left++

			if _, ok := need[ch]; ok {
				if window[ch] == need[ch] {
					valid--
				}
				window[ch]--
			}
		}
	}
	if minLength == math.MaxInt16 {
		return ""
	}

	return s[start: start+minLength]
}
