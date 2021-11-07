package leetcode0437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return pathWithRoot(root, targetSum) + pathWithRoot(root.Left, targetSum) + pathWithRoot(root.Right, targetSum)
}

func pathWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0
	if root.Val == sum {
		cnt++
	}

	cnt += pathWithRoot(root.Left, sum-root.Val)
	cnt += pathWithRoot(root.Right, sum-root.Val)

	return cnt
}
