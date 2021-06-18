package leetcode_0137

import "unsafe"

func singleNumber(nums []int) int {
	bits := int(unsafe.Sizeof(0) * 8)
	count := make([]int, bits)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < bits; j++ {
			count[j] += (nums[i] >> j) & 1
			count[j] %= 3
		}
	}

	ans := 0
	for i := 0; i < bits; i++ {
		ans += count[i] << i
	}
	return ans
}
