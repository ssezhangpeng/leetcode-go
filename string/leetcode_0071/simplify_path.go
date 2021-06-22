package leetcode_0071

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"strings"
)

func simplifyPath(path string) string {
	stk := linkedliststack.New()

	symbols := strings.Split(path, "/")

	for i := 0; i < len(symbols); i++ {
		if symbols[i] != "" && symbols[i] != "." {
			if symbols[i] == ".." {
				if !stk.Empty() {
					stk.Pop()
				}
			} else {
				stk.Push(symbols[i])
			}
		}
	}

	ans := ""
	for !stk.Empty() {
		val, _ := stk.Pop()
		ans = "/" + val.(string) + ans
	}
	if ans == "" {
		return "/"
	}
	return ans
}
