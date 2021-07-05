package leetcode_0108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTCore(nums, 0, len(nums)-1)
}

func sortedArrayToBSTCore(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	// make sure mid as greater value when [start, end]'s length is even
	mid := (start + end + 1) >> 1
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBSTCore(nums, start, mid-1)
	root.Right = sortedArrayToBSTCore(nums, mid+1, end)

	return root
}
