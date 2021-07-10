package leetcode_0041

func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		// nums[i] 应该放在 nums[i]-1 的下标处
		for nums[i] > 0 && nums[i] <= size && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < size; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return size + 1
}
