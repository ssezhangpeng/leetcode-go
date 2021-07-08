package leetcode_0129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	curr, sum := 0, 0
	sumNumbersCore(root, &curr, &sum)
	return sum
}

func sumNumbersCore(root *TreeNode, curr, sum *int) {
	if root == nil {
		return
	}

	*curr = 10**curr + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += *curr
	}

	sumNumbersCore(root.Left, curr, sum)
	sumNumbersCore(root.Right, curr, sum)

	*curr = (*curr - root.Val) / 10
}
