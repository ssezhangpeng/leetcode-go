package leetcode_0234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := getBeforeMid(head)

	// 断链
	afterHeader := mid.Next
	mid.Next = nil

	// 对后面的链表进行翻转
	afterHeader = reverse(afterHeader)

	// 判断是否是回文链表
	for head != nil && afterHeader != nil {
		if head.Val != afterHeader.Val {
			return false
		}
		head = head.Next
		afterHeader = afterHeader.Next
	}

	return true
}

func getBeforeMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
