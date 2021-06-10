package leetcode_0082

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{Val: -1, Next: nil}
	pre := &dummy
	pre.Next = head

	for pre.Next != nil {
		slow := pre.Next
		fast := slow.Next
		for fast != nil && slow.Val == fast.Val {
			fast = fast.Next
		}
		if fast != slow.Next {
			pre.Next = fast
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
