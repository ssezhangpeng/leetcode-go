package leetcode_0046

func permute(nums []int) [][]int {
	var ans [][]int
	var path []int
	visited := make([]bool, len(nums))

	backtracking(nums, path, &ans, visited)

	return ans
}

func backtracking(nums []int, path []int, ans *[][]int, visited []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		path = append(path, nums[i])
		visited[i] = true

		backtracking(nums, path, ans, visited)

		path = path[:len(path)-1]
		visited[i] = false
	}
}
