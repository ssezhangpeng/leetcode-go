package leetcode_0089

func grayCode(n int) []int {
	var size uint = 1 << uint(n)
	ans := make([]int, size)

	for i := uint(0); i < size; i++ {
		ans[i] = int(i ^ (i >> 1))
	}
	return ans
}
