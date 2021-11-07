package leetcode0543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	width := 1

	getTreeDepth(root, &width)

	// 最大宽度为最长路径上的节点数减一
	return width - 1
}

func getTreeDepth(root *TreeNode, width *int) int {
	if root == nil {
		return 0
	}

	leftHeight := getTreeDepth(root.Left, width)
	rightHeight := getTreeDepth(root.Right, width)

	// width 为以 root 节点作为根节点的最大路径上的节点数
	*width = max(*width, leftHeight+rightHeight+1)

	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
