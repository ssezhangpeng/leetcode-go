package leetcode_0102

func levelOrderRecursion(root *TreeNode) [][]int {
	var ans [][]int
	levelOrderRecursionCore(root, 1, &ans)
	return ans
}

func levelOrderRecursionCore(root *TreeNode, level int, ans *[][]int) {
	if root == nil {
		return
	}

	if level > len(*ans) {
		*ans = append(*ans, []int{})
	}
	(*ans)[level-1] = append((*ans)[level-1], root.Val)
	levelOrderRecursionCore(root.Left, level+1, ans)
	levelOrderRecursionCore(root.Right, level+1, ans)
}
