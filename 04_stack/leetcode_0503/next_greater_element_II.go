package leetcode_0503

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	var stk []int
	ans := make([]int, size)

	for i := 2*size - 1; i >= 0; i-- {
		for len(stk) > 0 && nums[i%size] >= stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) == 0 {
			ans[i%size] = -1
		} else {
			ans[i%size] = stk[len(stk)-1]
		}

		stk = append(stk, nums[i%size])
	}

	return ans
}
