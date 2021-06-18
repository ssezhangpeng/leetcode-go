package leetcode_0142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break
		}
	}

	slow = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
