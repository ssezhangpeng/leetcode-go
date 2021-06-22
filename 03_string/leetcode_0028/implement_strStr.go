package leetcode_0028

import "strings"

func strStr(haystack string, needle string) int {
	// KMP algorithm
	// Rabin-Karp algorithm in function of strings.Index()
	return strings.Index(haystack, needle)
}
