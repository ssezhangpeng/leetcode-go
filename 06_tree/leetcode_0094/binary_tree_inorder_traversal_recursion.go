package leetcode_0094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	inorderTraversalRecursion(root, &ans)
	return ans
}

func inorderTraversalRecursion(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	inorderTraversalRecursion(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorderTraversalRecursion(root.Right, ans)
}
