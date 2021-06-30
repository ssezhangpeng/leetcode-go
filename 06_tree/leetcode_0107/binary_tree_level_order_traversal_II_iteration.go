package leetcode_0107

func levelOrderBottomIteration(root *TreeNode) [][]int {
	var ans [][]int
	levelOrderBottomIterationCore(root, &ans)
	reverse(&ans)
	return ans
}

func levelOrderBottomIterationCore(root *TreeNode, ans *[][]int) {
	if root == nil {
		return
	}
	var current, next []*TreeNode
	current = append(current, root)

	for len(current) > 0 {
		var level []int
		for len(current) > 0 {
			node := current[0]
			current = current[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		*ans = append(*ans, level)
		current, next = next, current
	}
}
