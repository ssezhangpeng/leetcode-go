package leetcode_0109

import "math/bits"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	return sortedListToBSTCore(head)
}

func sortedListToBSTCore(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	preMid := getPreMid(head)
	mid := preMid.Next
	newHead := mid.Next

	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBSTCore(head)
	root.Right = sortedListToBSTCore(newHead)

	return root
}

func getPreMid(head *ListNode) *ListNode {
	return nil
}
