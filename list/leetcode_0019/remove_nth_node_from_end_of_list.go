package leetcode_0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{Val: -1, Next: nil}
	pre := &dummy
	pre.Next = head

	slow, fast := pre, pre
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
