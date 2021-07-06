package leetcode_0109

import (
	"fmt"
	"testing"
)

func generateList() *ListNode {
	var dummy ListNode
	pre := &dummy

	arr := []int{-10, -3, 0, 5, 9}

	for _, val := range arr {
		node := &ListNode{Val: val}
		pre.Next = node
		pre = pre.Next
	}
	return dummy.Next
}

func printTreeByLevel(root *TreeNode) {
	if root == nil {
		return
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d\t", node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func TestSortedListToBSTCore(t *testing.T) {
	head := generateList()

	root := sortedListToBSTCore(head)

	printTreeByLevel(root)
}
