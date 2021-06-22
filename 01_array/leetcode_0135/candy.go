package leetcode_0135

func candy(ratings []int) int {
	len := len(ratings)
	candy := make([]int, len)
	// init
	for i := 0; i < len; i++ {
		candy[i] = 1
	}
	// from left --> right
	for i := 0; i < len-1; i++ {
		if ratings[i] < ratings[i+1] {
			candy[i+1] = candy[i] + 1
		}
	}

	// from right --> left
	for i := len - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
	}
	ans := 0
	for i := 0; i < len; i++ {
		ans += candy[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
