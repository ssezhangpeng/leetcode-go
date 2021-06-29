package leetcode_0144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	preorderTraversalRecursion(root, &ans)
	return ans
}

func preorderTraversalRecursion(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	preorderTraversalRecursion(root.Left, ans)
	preorderTraversalRecursion(root.Right, ans)
}
