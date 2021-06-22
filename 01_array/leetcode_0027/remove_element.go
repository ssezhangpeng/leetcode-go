package leetcode_0027

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos
}
