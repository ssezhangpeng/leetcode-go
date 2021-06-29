package leetcode_1019

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return getMax(arr)
}

func getMax(num []int) []int {
	var stk []int
	ans := make([]int, len(num))

	for i := len(num) - 1; i >= 0; i-- {
		for len(stk) > 0 && num[i] >= stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stk[len(stk)-1]
		}

		stk = append(stk, num[i])
	}

	return ans
}
