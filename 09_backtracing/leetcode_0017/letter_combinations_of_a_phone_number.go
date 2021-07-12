package leetcode_0017

var dict = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var ans []string
	var path string

	backtracking(digits, 0, path, &ans)
	return ans
}

func backtracking(digits string, index int, path string, ans *[]string) {
	if len(path) == len(digits) {
		*ans = append(*ans, path)
		return
	}

	digit := digits[index] - '0'
	letters := dict[digit]
	for i := 0; i < len(letters); i++ {
		path += string(letters[i])
		backtracking(digits, index+1, path, ans)
		path = path[:len(path)-1]
	}
}
