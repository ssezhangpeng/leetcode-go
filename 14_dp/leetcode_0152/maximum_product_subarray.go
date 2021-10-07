package leetcode_0152

func maxProduct(nums []int) int {
	size := len(nums)
	maxF := make([]int, size)
	minF := make([]int, size)

	maxVal := nums[0]
	maxF[0], minF[0] = nums[0], nums[0]
	for i := 1; i < size; i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(minF[i-1]*nums[i], nums[i]))
		minF[i] = min(maxF[i-1]*nums[i], min(minF[i-1]*nums[i], nums[i]))
		maxVal = max(maxVal, maxF[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
