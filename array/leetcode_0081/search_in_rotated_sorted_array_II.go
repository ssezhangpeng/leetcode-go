package leetcode_0081

// [1, 3, 1, 1, 1]
func search(nums []int, target int) bool {
	len := len(nums)

	l, r := 0, len-1
	for l <= r {
		mid := l + (r-l) >> 1
		if nums[mid] == target {
			return true
		}

		if nums[l] < nums[mid] { // [l, mid] asc
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[l] > nums[mid] { // [mid, r] asc
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			l++
		}
	}
	return false
}
