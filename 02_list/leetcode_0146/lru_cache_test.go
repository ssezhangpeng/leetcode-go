package leetcode_0146

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2) // nil
	obj.Put(1, 1)         // nil
	obj.Put(2, 2)         // nil
	g1 := obj.Get(1)
	fmt.Println("g1: ", g1) // 1

	obj.Put(3, 3) // nil
	g2 := obj.Get(2)
	fmt.Println("g2: ", g2) // -1

	obj.Put(4, 4)
	g1 = obj.Get(1)
	fmt.Println("g1: ", g1)

	g3 := obj.Get(3)
	fmt.Println("g3: ", g3)

	g4 := obj.Get(4)
	fmt.Println("g4: ", g4)
}
