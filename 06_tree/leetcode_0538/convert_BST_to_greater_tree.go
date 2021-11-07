package leetcode0538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	inorderReverse(root, &sum)

	return root
}

func inorderReverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	inorderReverse(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	inorderReverse(root.Left, sum)
}
