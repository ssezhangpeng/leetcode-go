package leetcode_0034

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)

	if first == -1 || last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l <= r {
		mid := l + (r - l) >> 1
		if nums[mid] == target {
			if mid == l || nums[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func findLast(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l <= r {
		mid := l + (r - l) >> 1
		if nums[mid] == target {
			if mid == r || nums[mid+1] != target {
				return mid
			} else {
				l = mid + 1
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
