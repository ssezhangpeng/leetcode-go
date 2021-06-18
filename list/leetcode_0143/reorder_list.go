package leetcode_0143

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	dummy := ListNode{Val: -1, Next: nil}
	dummy.Next = head
	prev := head

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// disconnect linklist
	lastHead := slow.Next
	slow.Next = nil

	print(lastHead)
	lastHead = reverse(lastHead)
	print(lastHead)

	for lastHead != nil {
		pNext := lastHead.Next
		lastHead.Next = prev.Next
		prev.Next = lastHead
		prev = prev.Next.Next
		lastHead = pNext
	}
	head = dummy.Next
}

func reverse(head *ListNode) *ListNode {
	dummy := ListNode{Val: -1, Next: nil}
	prev := &dummy

	for head != nil {
		pNext := head.Next
		head.Next = prev.Next
		prev.Next = head

		head = pNext
	}
	return dummy.Next
}

func print(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
