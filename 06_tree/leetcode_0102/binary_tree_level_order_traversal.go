package leetcode_0102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
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
		ans = append(ans, level)
		current, next = next, current
	}
	return ans
}
