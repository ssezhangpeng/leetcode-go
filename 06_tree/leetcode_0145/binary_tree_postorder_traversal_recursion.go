package leetcode_0145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	postorderTraversalRecursion(root, &ans)
	return ans
}

func postorderTraversalRecursion(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	postorderTraversalRecursion(root.Left, ans)
	postorderTraversalRecursion(root.Right, ans)
	*ans = append(*ans, root.Val)
}
