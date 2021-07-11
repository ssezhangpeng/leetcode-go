package leetcode_0077

func combine(n int, k int) [][]int {
	var ans [][]int
	var path []int

	backtracking(n, k, 1, path, &ans)

	return ans
}

func backtracking(n, k, startIndex int, path []int, ans *[][]int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
