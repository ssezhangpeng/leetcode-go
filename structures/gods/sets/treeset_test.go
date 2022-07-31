package sets

import (
	"testing"

	"github.com/emirpasic/gods/sets/treeset"
)

func TestMain(t *testing.T) {
	set := treeset.NewWithIntComparator() // empty (keys are of type int)
	set.Add(1)                            // 1
	set.Add(2, 2, 3, 4, 5)                // 1, 2, 3, 4, 5 (in order, duplicates ignored)

	// 可以取出所有元素，拿到有序列表中的第一个
	is := set.Values()
	if value, ok := is[0].(int); !ok {
		t.Fatal("It is not ok for type int.")
	} else {
		t.Logf("--->value: %d", value)
	}

	set.Remove(4)      // 1, 2, 3, 5 (in order)
	set.Remove(2, 3)   // 1, 5 (in order)
	set.Contains(1)    // true
	set.Contains(1, 5) // true
	set.Contains(1, 6) // false
	_ = set.Values()   // []int{1,5} (in order)
	set.Clear()        // empty
	set.Empty()        // true
	set.Size()         // 0
}
