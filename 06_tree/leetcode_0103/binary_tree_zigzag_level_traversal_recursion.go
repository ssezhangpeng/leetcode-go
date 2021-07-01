package leetcode_0103

func zigzagLevelOrderRecursion(root *TreeNode) [][]int {
	var ans [][]int
	zigzagLevelOrderRecursionCore(root, 1, &ans, false)
	return ans
}

func zigzagLevelOrderRecursionCore(root *TreeNode, level int, ans *[][]int, hasReverse bool) {
	if root == nil {
		return
	}

	if level > len(*ans) {
		*ans = append(*ans, []int{})
	}

	if hasReverse {
		(*ans)[level-1] = append((*ans)[level-1], 0)
		copy((*ans)[level-1][1:], (*ans)[level-1][0:])
		(*ans)[level-1][0] = root.Val
	} else {
		(*ans)[level-1] = append((*ans)[level-1], root.Val)
	}

	zigzagLevelOrderRecursionCore(root.Left, level+1, ans, !hasReverse)
	zigzagLevelOrderRecursionCore(root.Right, level+1, ans, !hasReverse)
}
