package leetcode_0035

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r + 1
}
