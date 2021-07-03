package leetcode_0106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	return buildTreeCore(inorder, 0, len(inorder), postorder, 0, len(inorder))
}

func buildTreeCore(inorder []int, iB, iE int, postorder []int, pB, pE int) *TreeNode {
	if iB == iE || pB == pE {
		return nil
	}

	root := &TreeNode{Val: postorder[pE-1]}

	pos := find(inorder, iB, iE, postorder[pE-1])
	leftTreeSize := pos - iB

	root.Left = buildTreeCore(inorder, iB, pos, postorder, pB, pB+leftTreeSize)
	root.Right = buildTreeCore(inorder, pos+1, iE, postorder, pB+leftTreeSize, pE-1)

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
