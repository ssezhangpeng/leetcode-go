package leetcode_0025

type ListNode struct {
	Val  int
	Next *ListNode
}

//TODO: not understand still now
func reverseKGroup(head *ListNode, k int) *ListNode {
	len := getLen(head)
	group := len / k

	dummy := ListNode{Val: -1, Next: nil}
	pre := &dummy
	pre.Next = head

	for i := 0; i < group; i++ {
		tail, newHead := getHeadTail(pre, k)
		for j := 0; j < k; j++ {
			pNext := head.Next
			head.Next = pre.Next
			pre.Next = head
			head = pNext
		}
		tail.Next = newHead
		head = newHead
		pre = tail
	}

	return dummy.Next
}

func getLen(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}

func getHeadTail(pre *ListNode, k int) (*ListNode, *ListNode) {
	first := pre.Next
	for i := 0; i < k; i++ {
		pre = pre.Next
	}
	return first, pre.Next
}
