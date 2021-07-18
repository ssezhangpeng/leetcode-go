package leetcode_0045

func jump(nums []int) int {
	step := 0
	left, right := 0, 0 // 每次可跳跃的覆盖范围：[left, right]

	if len(nums) == 1 {
		return step
	}

	for left <= right {
		step++
		oldRight := right

		for i := left; i <= oldRight; i++ {
			newRight := i + nums[i]
			if newRight >= len(nums)-1 {
				return step
			}
			if newRight > right {
				right = newRight
			}
		}
		left = oldRight + 1
	}
	return 0
}
