package leetcode_0012

func intToRoman(num int) string {
	radix := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; num > 0; i++ {
		count := num / radix[i]
		num %= radix[i]
		for ; count > 0; count-- {
			roman += symbol[i]
		}
	}
	return roman
}
