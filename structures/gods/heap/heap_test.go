package heap

import (
	"testing"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func TestMinHeap(t *testing.T) {
	// Min-heap
	heap := binaryheap.NewWithIntComparator() // empty (min-heap)
	heap.Push(2)                              // 2
	heap.Push(3)                              // 2, 3
	heap.Push(1)                              // 1, 3, 2
	heap.Values()                             // 1, 3, 2
	_, _ = heap.Peek()                        // 1,true
	_, _ = heap.Pop()                         // 1, true
	_, _ = heap.Pop()                         // 2, true
	_, _ = heap.Pop()                         // 3, true
	_, _ = heap.Pop()                         // nil, false (nothing to pop)
	heap.Push(1)                              // 1
	heap.Clear()                              // empty
	heap.Empty()                              // true
	heap.Size()                               // 0
}

func TestMaxHeap(t *testing.T) {
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	heap := binaryheap.NewWith(inverseIntComparator) // empty (max-heap)

	heap.Push(2)       // 2
	heap.Push(3)       // 2, 3
	heap.Push(1)       // 1, 3, 2
	heap.Values()      // 1, 3, 2
	_, _ = heap.Peek() // 1,true
	_, _ = heap.Pop()  // 1, true
	_, _ = heap.Pop()  // 2, true
	_, _ = heap.Pop()  // 3, true
	_, _ = heap.Pop()  // nil, false (nothing to pop)
	heap.Push(1)       // 1
	heap.Clear()       // empty
	heap.Empty()       // true
	heap.Size()        // 0
}
