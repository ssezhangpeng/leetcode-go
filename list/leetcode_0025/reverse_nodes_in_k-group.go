package leetcode_0025

type ListNode struct {
	Val int
	Next *ListNode
}
//TODO: not understand still now
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

func getLen(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}
