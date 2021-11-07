package leetcode0494

func findTargetSumWays(nums []int, target int) int {
	cnt := 0
	backtrack(nums, 0, 0, target, &cnt)
	return cnt
}

func backtrack(nums []int, start int, sum int, target int, cnt *int) {
	if start == len(nums) {
		if sum == target {
			(*cnt)++
		}
		return
	}

	// 将 start 下标的元素正数加入
	backtrack(nums, start+1, sum+nums[start], target, cnt)
	// 将 start 下标的元素负数加入
	backtrack(nums, start+1, sum-nums[start], target, cnt)
}
