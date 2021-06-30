package leetcode_0145

func postorderTraversalIteration(root *TreeNode) []int {
	// mid-->right-->left
	ans := postorderTraversalIterationCore(root)
	// left-->right-->mid
	reverse(&ans)
	return ans
}

func postorderTraversalIterationCore(root *TreeNode) []int {
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

		if node.Left != nil {
			stk = append(stk, node.Left)
		}

		if node.Right != nil {
			stk = append(stk, node.Right)
		}
	}
	return ans
}

func reverse(num *[]int) {
	for i, j := 0, len(*num)-1; i < j; i, j = i+1, j-1 {
		(*num)[i], (*num)[j] = (*num)[j], (*num)[i]
	}
	return
}
