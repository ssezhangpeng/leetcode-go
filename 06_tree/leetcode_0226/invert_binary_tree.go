package leetcode_0226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// 先序遍历进行翻转
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}
