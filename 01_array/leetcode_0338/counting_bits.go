package leetcode_0338

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = countOnes(i)
	}

	return ans
}

func countOnes(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}

	return count
}
