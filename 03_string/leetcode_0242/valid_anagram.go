package leetcode_0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cntMap := map[rune]int{}
	for _, ch := range s {
		cntMap[ch]++
	}

	for _, ch := range t {
		cntMap[ch]--
		if cntMap[ch] < 0 {
			return false
		}
	}

	return true
}
