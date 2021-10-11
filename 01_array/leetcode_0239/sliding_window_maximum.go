package leetcode_0239

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)

	// window 队列中元素始终保持递减排序，window[0] 始终是窗口内的最大元素
	// 如果后面的元素比较大，则前面相对小的元素永远不可能成为最大元素，故需要一直出队
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(window) > 0 && nums[i] > nums[window[len(window)-1]] {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		if i-window[0]+1 > k {
			window = window[1:]
		}

		if i >= k-1 {
			ans = append(ans, nums[window[0]])
		}
	}

	return ans
}
