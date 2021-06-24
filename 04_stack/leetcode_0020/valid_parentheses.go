package leetcode_0020

import "github.com/emirpasic/gods/stacks/linkedliststack"

func isValid(s string) bool {
	stk := linkedliststack.New()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk.Push(s[i])
		} else {
			if stk.Empty() {
				return false
			}
			topVal, _ := stk.Peek()
			if s[i] == ')' && topVal.(byte) == '(' ||
				s[i] == '}' && topVal.(byte) == '{' ||
				s[i] == ']' && topVal.(byte) == '[' {
				stk.Pop()
			} else {
				return false
			}
		}
	}
	return stk.Empty()
}
