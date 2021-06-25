package leetcode_0402

func removeKdigits(num string, k int) string {
	var stkAsc []int

	for i := 0; i < len(num); i++ {
		for len(stkAsc) > 0 && num[i] < num[stkAsc[len(stkAsc)-1]] && k > 0 {
			stkAsc = stkAsc[:len(stkAsc)-1]
			k--
		}
		stkAsc = append(stkAsc, i)
	}

	for k > 0 {
		stkAsc = stkAsc[:len(stkAsc)-1]
		k--
	}

	ans, idx := "", 0

	for ; idx < len(stkAsc); idx++ {
		if num[stkAsc[idx]] != '0' {
			break
		}
	}

	for ; idx < len(stkAsc); idx++ {
		ans += string(num[stkAsc[idx]])
	}

	if ans == "" {
		return "0"
	}

	return ans
}
