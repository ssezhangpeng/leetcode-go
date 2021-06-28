package leetcode_0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stk []int
	m := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stk) > 0 && nums2[i] >= stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stk[len(stk)-1]
		}

		stk = append(stk, nums2[i])
	}

	var ans []int
	for i := 0; i < len(nums1); i++ {
		nextVal, _ := m[nums1[i]]
		ans = append(ans, nextVal)
	}
	return ans
}
