package leetcode_0239

// TODO: 未完成
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)

	// window 中元素始终保持递减排序，window[0] 始终是窗口内的最大元素
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(window) > 0 && nums[i] > nums[window[0]] {
			window = window[1:]
		}

		pushFront(window, i)

		if i-window[0]+1 > k {
			window = window[1:]
		}

		if i >= k-1 {
			ans = append(ans, nums[window[0]])
		}
	}

	return ans
}

func pushFront(arr []int, target int) {
	temp := make([]int, len(arr)+1)
	copy(temp[1:], arr)
	temp[0] = target

	arr = temp
	return
}
