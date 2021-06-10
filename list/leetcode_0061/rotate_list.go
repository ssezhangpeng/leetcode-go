package leetcode_0061

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	len := 1
	p := head
	for p.Next != nil {
		len++
		p = p.Next
	}
	k %= len
	p.Next = head // link to circle

	p = head
	dist := len - k
	for i:=0; i<dist-1; i++ {
		p = p.Next
	}

	head = p.Next
	p.Next = nil

	return head
}
