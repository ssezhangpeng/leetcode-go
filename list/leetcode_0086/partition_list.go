package leetcode_0086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var lessHead, moreHead ListNode
	lessPtr, morePtr := &lessHead, &moreHead

	for head != nil {
		if head.Val < x {
			lessPtr.Next = head
			lessPtr = lessPtr.Next
		} else {
			morePtr.Next = head
			morePtr = morePtr.Next
		}
		head = head.Next
	}
	morePtr.Next = nil
	lessPtr.Next = moreHead.Next
	return lessHead.Next
}
