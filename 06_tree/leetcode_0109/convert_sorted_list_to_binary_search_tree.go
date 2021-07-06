package leetcode_0109

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

	pre, mid := getPreMid(head)

	if pre == mid {
		return &TreeNode{Val: mid.Val}
	}

	pre.Next = nil
	newHead := mid.Next
	mid.Next = nil

	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBSTCore(head)
	root.Right = sortedListToBSTCore(newHead)

	return root
}

func getPreMid(head *ListNode) (pre *ListNode, mid *ListNode) {
	if head == nil {
		return nil, nil
	}

	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre, slow
}
