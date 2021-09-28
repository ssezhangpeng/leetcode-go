package leetcode_0155

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println("min: ", obj.GetMin())
	obj.Pop()
	fmt.Println("top: ", obj.Top())
	fmt.Println("min: ", obj.GetMin())
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */