package leetcode_0107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	levelOrderBottomRecursion(root, 1, &ans)
	reverse(&ans)
	return ans
}

func levelOrderBottomRecursion(root *TreeNode, level int, ans *[][]int) {
	if root == nil {
		return
	}

	if level > len(*ans) {
		*ans = append(*ans, []int{})
	}

	(*ans)[level-1] = append((*ans)[level-1], root.Val)
	levelOrderBottomRecursion(root.Left, level+1, ans)
	levelOrderBottomRecursion(root.Right, level+1, ans)
}

func reverse(num *[][]int) {
	for i, j := 0, len(*num)-1; i < j; i, j = i+1, j-1 {
		(*num)[i], (*num)[j] = (*num)[j], (*num)[i]
	}
}
