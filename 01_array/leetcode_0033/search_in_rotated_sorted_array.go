package leetcode_0033

func search(nums []int, target int) int {
	len := len(nums)
	l, r := 0, len-1

	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[l] { //[l, mid] asc
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // [mid, r] asc
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
