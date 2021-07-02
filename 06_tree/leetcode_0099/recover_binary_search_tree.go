package leetcode_0099

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var pre, target1, target2 *TreeNode
	var stk []*TreeNode

	for len(stk) > 0 || root != nil {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}

		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if pre != nil && pre.Val > node.Val {
			target2 = node
			if target1 == nil {
				target1 = pre
			} else {
				break
			}
		}
		pre = node
		root = node.Right
	}
	target1.Val, target2.Val = target2.Val, target1.Val
}
