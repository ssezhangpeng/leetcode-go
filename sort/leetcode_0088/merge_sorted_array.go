package leetcode_0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		nums[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, nums)
}
