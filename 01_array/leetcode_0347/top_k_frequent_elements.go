package leetcode0347

import (
	"github.com/emirpasic/gods/trees/binaryheap"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	// 根据 count 建立小根堆，堆中元素为 key
	comparator := func(a, b interface{}) int {
		return m[a.(int)] - m[b.(int)]
	}
	minHeap := binaryheap.NewWith(comparator)
	for key, count := range m {
		if minHeap.Size() < k {
			minHeap.Push(key)
		} else {
			topCount, _ := minHeap.Peek()
			if count > m[topCount.(int)] {
				minHeap.Pop()
				minHeap.Push(key)
			}
		}
	}

	// 取出堆中的k个元素
	var ans []int
	for minHeap.Size() > 0 {
		value, _ := minHeap.Pop()
		ans = append(ans, value.(int))
	}

	return ans
}
