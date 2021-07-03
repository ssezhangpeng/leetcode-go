package leetcode_0117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	p := root

	// connect every level
	for p != nil {
		var dummy Node
		var prev = &dummy
		// connect next one level
		for p != nil {
			if p.Left != nil {
				prev.Next = p.Left
				prev = prev.Next
			}

			if p.Right != nil {
				prev.Next = p.Right
				prev = prev.Next
			}

			p = p.Next
		}
		// set p as header at next level
		p = dummy.Next
	}
	return root
}
