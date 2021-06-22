package leetcode_0083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = nil
	return head
}
