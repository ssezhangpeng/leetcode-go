package leetcode_0024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{Val: -1, Next: nil}
	pre := &dummy
	pre.Next = head

	for pre.Next != nil && pre.Next.Next != nil {
		slow, fast := pre.Next, pre.Next.Next
		pNext := fast.Next

		pre.Next = fast
		fast.Next = slow
		slow.Next = pNext

		pre = pre.Next.Next
	}

	return dummy.Next
}
