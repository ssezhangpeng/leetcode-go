package leetcode_0101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val &&
		isSymmetricCore(p.Left, q.Right) &&
		isSymmetricCore(p.Right, q.Left)
}
