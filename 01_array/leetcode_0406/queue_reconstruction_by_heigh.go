package leetcode0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// step1: 进行身高从高到低排序
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})

	// step2: 遍历后面的元素，按 K 值进行放置
	var ans [][]int
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}

	return ans
}
