package leetcode_0098

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidBSTCore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCore(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return isValidBSTCore(root.Left, lower, root.Val) &&
		isValidBSTCore(root.Right, root.Val, upper)
}
