package leetcode_0321

// ref: https://www.bilibili.com/video/BV1MA411T7YR?p=4

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	len1, len2 := len(nums1), len(nums2)

	var preMax []int
	for i := 0; i <= k; i++ {
		if i > len1 || k-i > len2 {
			continue
		}

		max1 := getMax(nums1, i)
		max2 := getMax(nums2, k-i)

		temp := merge(max1, max2)
		if greater(temp, preMax, 0, 0) {
			preMax = temp
		}
	}
	return preMax
}

func getMax(nums []int, k int) []int {
	var descStk []int
	delCnt := len(nums) - k

	for i := 0; i < len(nums); i++ {
		for len(descStk) > 0 && nums[i] > descStk[len(descStk)-1] && delCnt > 0 {
			descStk = descStk[:len(descStk)-1]
			delCnt--
		}
		descStk = append(descStk, nums[i])
	}
	return descStk[:k]
}

func merge(num1, num2 []int) []int {
	var ans []int
	len1, len2 := len(num1), len(num2)
	index1, index2 := 0, 0

	for index1 < len1 && index2 < len2 {
		if greater(num1, num2, index1, index2) {
			ans = append(ans, num1[index1])
			index1++
		} else {
			ans = append(ans, num2[index2])
			index2++
		}
	}

	for index1 < len1 {
		ans = append(ans, num1[index1])
		index1++
	}
	for index2 < len2 {
		ans = append(ans, num2[index2])
		index2++
	}

	return ans
}

func greater(num1, num2 []int, index1, index2 int) bool {
	len1, len2 := len(num1), len(num2)

	for index1 < len1 && index2 < len2 {
		if num1[index1] > num2[index2] {
			return true
		} else if num1[index1] < num2[index2] {
			return false
		} else {
			index1++
			index2++
		}
	}
	return index1 != len1
}
