package leetcode_0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stkNumber []string

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			val2, _ := strconv.Atoi(stkNumber[len(stkNumber)-1])
			val1, _ := strconv.Atoi(stkNumber[len(stkNumber)-2])
			stkNumber = stkNumber[:len(stkNumber)-2]
			val := calculate(val1, val2, tokens[i])
			stkNumber = append(stkNumber, strconv.Itoa(val))
		} else {
			stkNumber = append(stkNumber, tokens[i])
		}
	}
	ans, _ := strconv.Atoi(stkNumber[0])
	return ans
}

func calculate(a, b int, op string) int {
	ans := 0
	switch op {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	default:
		panic("invalid op")
	}
	return ans
}
