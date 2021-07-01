package leetcode_0103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	hasReverse := false
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
		if hasReverse {
			reverse(&level)
		}
		ans = append(ans, level)

		hasReverse = !hasReverse
		current, next = next, current
	}
	return ans
}

func reverse(nums *[]int) {
	for i, j := 0, len(*nums)-1; i < j; i, j = i+1, j-1 {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
}
