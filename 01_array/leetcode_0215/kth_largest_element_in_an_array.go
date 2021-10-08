package leetcode_0215

import (
	"github.com/emirpasic/gods/trees/binaryheap"
)

func findKthLargest(nums []int, k int) int {
	minHeap := binaryheap.NewWithIntComparator()

	for i := 0; i < len(nums); i++ {
		if minHeap.Size() < k {
			minHeap.Push(nums[i])
		} else {
			value, _ := minHeap.Peek()
			topValue, _ := value.(int)
			if topValue > nums[i] {
				continue
			} else {
				minHeap.Pop()
				minHeap.Push(nums[i])
			}
		}
	}

	topValue, _ := minHeap.Peek()
	return topValue.(int)
}
