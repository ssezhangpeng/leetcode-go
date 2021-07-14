package leetcode_0131

func partition(s string) [][]string {
	var ans [][]string
	var path []string

	backtracking(s, 0, path, &ans)

	return ans
}

func backtracking(s string, start int, path []string, ans *[][]string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			backtracking(s, i+1, path, ans)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end && s[start] == s[end] {
		start++
		end--
	}
	return start >= end
}
