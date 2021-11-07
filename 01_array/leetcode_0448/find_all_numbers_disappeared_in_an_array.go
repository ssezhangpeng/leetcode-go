package leetcode0448

func findDisappearedNumbers(nums []int) []int {
	size := len(nums)

	for i := 0; i < size; i++ {
		idx := (nums[i] - 1) % size
		nums[idx] += size
	}

	var result []int
	// idx 处应该放置的元素是 idx+1
	for idx := 0; idx < size; idx++ {
		if nums[idx] <= size {
			result = append(result, idx+1)
		}
	}

	return result
}
