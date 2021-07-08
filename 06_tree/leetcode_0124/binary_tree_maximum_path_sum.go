package leetcode_0124

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxVal := math.MinInt8
	maxPathSumCore(root, &maxVal)
	return maxVal
}

// 以 root 为根节点的树的最大路径和
func maxPathSumCore(root *TreeNode, maxVal *int) int {
	if root == nil {
		return math.MinInt8
	}

	left := maxPathSumCore(root.Left, maxVal)
	right := maxPathSumCore(root.Right, maxVal)

	rootMax := max(max(left, right)+root.Val, root.Val)
	*maxVal = max(*maxVal, max(rootMax, left+right+root.Val))

	return rootMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
