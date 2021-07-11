package leetcode_0078

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int

	backTrace(nums, 0, path, &ans)

	return ans
}

func backTrace(nums []int, startIndex int, path []int, ans *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ans = append(*ans, temp)

	// i 为横向搜索分支数
	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		backTrace(nums, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
