package leetcode_0031

func nextPermutation(nums []int) {
	len := len(nums)
	if len < 2 {
		return
	}

	i, j := 0, 0
	// step1: find arr[i] < arr[i+1], from right --> left
	for i = len - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums, 0, len-1)
		return
	}
	// step2: find arr[j] > arr[i], from right --> left
	for j = len - 1; j > i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}

	// step3: swap(arr[i], arr[j])
	nums[i], nums[j] = nums[j], nums[i]
	// step4: reverse(arr[i+1], arr[len])
	reverse(nums, i+1, len-1)
}

func reverse(num []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
