package leetcode_0021

import (
	"fmt"
	"testing"
)

func generateList() *ListNode {
	arr := []int{-1, 5, 3, 4, 0}

	var dummy ListNode
	pre := &dummy

	for i := 0; i < len(arr); i++ {
		node := &ListNode{Val: arr[i]}
		pre.Next = node
		pre = pre.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func TestInsertionSortList(t *testing.T) {
	head := generateList()
	printList(head)
	newHead := insertionSortList(head)
	printList(newHead)
}
