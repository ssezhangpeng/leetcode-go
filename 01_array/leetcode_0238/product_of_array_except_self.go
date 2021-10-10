package leetcode_0238

func productExceptSelf(nums []int) []int {
	size := len(nums)

	left := make([]int, size)
	right := make([]int, size)

	left[0], right[size-1] = 1, 1
	// left[i] 标识下标 i 左边所有元素的乘积
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// right[i] 标识下标 i 右边所有元素的乘积
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// 获取结果
	ans := make([]int, size)
	for i := 0; i < size; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
