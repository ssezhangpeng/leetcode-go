package leetcode_0206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func reverseList(head *ListNode) *ListNode {
	// 递归出口
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
