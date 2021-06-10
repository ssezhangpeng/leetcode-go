package leetcode_0092

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummy ListNode
	pre := &dummy
	pre.Next = head
	// find left Node && prev left Node
	var prevLeft *ListNode
	l := pre
	for left > 0 {
		prevLeft = l
		l = l.Next
		left--
	}
	// find right Node
	var nextRight *ListNode
	r := pre
	for right > 0 {
		r = r.Next
		right--
	}
	nextRight = r.Next
	// parser linked list
	prevLeft.Next = nil
	r.Next = nil

	// reverse subList
	subListHead, subListTail := reverse(l)

	prevLeft.Next = subListHead
	subListTail.Next = nextRight
	return dummy.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode){
	var dummy ListNode
	pre := &dummy

	var nextNode *ListNode
	for head != nil {
		nextNode = head.Next
		head.Next = pre.Next
		pre.Next = head

		head = nextNode
	}

	// find tail
	tail := &dummy
	for tail.Next != nil {
		tail = tail.Next
	}

	return dummy.Next, tail
}
