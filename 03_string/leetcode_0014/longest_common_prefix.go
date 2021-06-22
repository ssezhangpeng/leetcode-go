package leetcode_0014

import "sort"

type ByLength []string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// put shortest string on the first position
	sort.Sort(ByLength(strs))

	for idx := 0; idx < len(strs[0]); idx++ {
		for j := 1; j < len(strs); j++ {
			str := strs[j]
			if str[idx] != strs[0][idx] {
				return strs[0][0:idx]
			}
		}
	}
	return strs[0]
}

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
