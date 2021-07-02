package leetcode_0110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return isBalancedCore(root) >= 0
}

func isBalancedCore(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	lHeight := isBalancedCore(root.Left)
	rHeight := isBalancedCore(root.Right)

	if lHeight < 0 || rHeight < 0 || math.Abs(lHeight-rHeight) > 1 {
		return -1
	}

	return math.Max(lHeight, rHeight) + 1
}
