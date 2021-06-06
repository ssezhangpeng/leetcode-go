package leetcode_0977

func sortedSquares(nums []int) []int {
	length := len(nums)

	l, r := 0, length-1
	pos := length - 1
	ans := make([]int, length)

	for l <= r {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			ans[pos] = nums[l] * nums[l]
			l++
		} else {
			ans[pos] = nums[r] * nums[r]
			r--
		}
		pos--
	}

	return ans
}
