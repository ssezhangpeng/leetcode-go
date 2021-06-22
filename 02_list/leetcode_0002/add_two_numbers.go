package leetcode_0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy ListNode
	pre := &dummy
	carry := 0
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		val := (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10
		node := &ListNode{Val: val, Next: nil}
		pre.Next = node
		pre = pre.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		pre.Next = &ListNode{Val: carry, Next: nil}
	}
	return dummy.Next
}
