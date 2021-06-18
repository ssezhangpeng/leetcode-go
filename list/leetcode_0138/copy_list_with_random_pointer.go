package leetcode_0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	dummy := Node{Val: -1, Next: nil, Random: nil}
	pre := &dummy

	// copy value
	p := head
	for p != nil {
		node := &Node{Val: p.Val, Next: nil, Random: nil}
		m[p] = node
		pre.Next = node
		pre = pre.Next
		p = p.Next
	}
	// copy randomPtr
	p, q := head, dummy.Next
	for p != nil {
		q.Random = m[p.Random]
		p = p.Next
		q = q.Next
	}

	return dummy.Next
}
