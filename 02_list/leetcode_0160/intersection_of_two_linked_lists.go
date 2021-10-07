package leetcode_0160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA := getListLength(headA)
	lenB := getListLength(headB)

	diff := lenA - lenB
	if lenA < lenB {
		diff = lenB - lenA
	}

	pA, pB := headA, headB
	// 更长的链表先走 diff 步
	for diff > 0 {
		if lenA < lenB {
			pB = pB.Next
		} else {
			pA = pA.Next
		}
		diff--
	}

	// 同时向后移动，直到链表末尾
	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}

	return nil
}

func getListLength(head *ListNode) int {
	length := 0

	for head != nil {
		length++
		head = head.Next
	}

	return length
}
