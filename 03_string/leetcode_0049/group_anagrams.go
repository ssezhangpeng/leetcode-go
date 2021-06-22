package leetcode_0049

import (
	"sort"
)

type ByByte []byte

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	var ans [][]string

	for i := 0; i < len(strs); i++ {
		bs := ByByte(strs[i])
		sort.Sort(bs)

		key := string(bs)
		m[key] = append(m[key], strs[i])
	}

	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func (b ByByte) Len() int {
	return len(b)
}

func (b ByByte) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b ByByte) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
