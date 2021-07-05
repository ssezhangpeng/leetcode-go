package leetcode_0095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateBSTrees(1, n)
}

func generateBSTrees(start, end int) []*TreeNode {
	tree := []*TreeNode{}

	if start > end {
		tree = append(tree, nil)
		return tree
	}

	for i := start; i <= end; i++ {
		left := generateBSTrees(start, i-1)
		right := generateBSTrees(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
