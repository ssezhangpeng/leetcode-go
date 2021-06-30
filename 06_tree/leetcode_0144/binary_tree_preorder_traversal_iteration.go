package leetcode_0144

func preorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	var stk []*TreeNode

	stk = append(stk, root)
	for len(stk) > 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		ans = append(ans, node.Val)

		if node.Right != nil {
			stk = append(stk, node.Right)
		}
		if node.Left != nil {
			stk = append(stk, node.Left)
		}
	}
	return ans
}
