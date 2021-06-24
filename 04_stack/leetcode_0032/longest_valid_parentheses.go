package leetcode_0032

import "github.com/emirpasic/gods/stacks/linkedliststack"

func longestValidParentheses(s string) int {
	maxLen, last := 0, -1        // the position of the last ')'
	stk := linkedliststack.New() // keep track of the position of mismatch '('

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk.Push(i)
		} else {
			if stk.Empty() {
				last = i
			} else {
				stk.Pop()
				if stk.Empty() {
					maxLen = max(maxLen, i-last)
				} else {
					left, _ := stk.Peek()
					maxLen = max(maxLen, i-left.(int))
				}
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
