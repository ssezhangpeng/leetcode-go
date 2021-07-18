package leetcode_0055

func canJump(nums []int) bool {
	reachIndex := 0
	for i := 0; i <= reachIndex && reachIndex < len(nums); i++ {
		reachIndex = max(reachIndex, i+nums[i])
	}
	return reachIndex >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
