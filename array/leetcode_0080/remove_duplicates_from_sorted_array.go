package leetcode_0080

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len < 2 {
		return len
	}

	pos, dups := 1, 1
	for j := 1; j < len; j++ {
		if nums[j] == nums[j-1] && dups+1 <= 2 {
			dups++
			nums[pos] = nums[j]
			pos++
		} else if nums[j] != nums[j-1] {
			dups = 1
			nums[pos] = nums[j]
			pos++
		}
	}
	return pos
}
