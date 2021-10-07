package leetcode_0169

func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count = 1
		} else if nums[i] == majority {
			count++
		} else {
			count--
		}
	}

	// 判断 majority 出现次数是不是超半, 题目已经假设: 给定的数组总是存在多数元素
	return majority
}
