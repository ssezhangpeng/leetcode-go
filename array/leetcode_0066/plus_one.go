package leetcode_0066

func plusOne(digits []int) []int {
	return plusOneCore(digits, 1)
}

func plusOneCore(digits []int, carry int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}
	if carry > 0 {
		digits = append(digits, -1)
		copy(digits[1:], digits[0:])
		digits[0] = carry
	}
	return digits
}
