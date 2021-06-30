package leetcode_0094

func inorderTraversalIteration(root *TreeNode) []int {
	var ans []int
	var stk []*TreeNode

	for len(stk) > 0 || root != nil {
		if root != nil {
			stk = append(stk, root)
			root = root.Left
		} else {
			node := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			ans = append(ans, node.Val)
			root = node.Right
		}
	}

	return ans
}
