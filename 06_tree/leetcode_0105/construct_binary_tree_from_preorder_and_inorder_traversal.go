package leetcode_0105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	return buildTreeCore(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func buildTreeCore(preorder []int, pB, pE int, inorder []int, iB, iE int) *TreeNode {
	if pB == pE || iB == iE {
		return nil
	}

	root := &TreeNode{Val: preorder[pB]}

	pos := find(inorder, iB, iE, preorder[pB])
	leftTreeSize := pos - iB

	root.Left = buildTreeCore(preorder, pB+1, pB+leftTreeSize+1, inorder, iB, pos)
	root.Right = buildTreeCore(preorder, pB+leftTreeSize+1, pE, inorder, pos+1, iE)

	return root
}

func find(arr []int, start, end, val int) int {
	for i := start; i < end; i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}
