package leetcode_0321

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := []int{5, 5, 1}
	k := 2

	ans := getMax(arr, k)
	fmt.Printf("ans: %+v\n", ans)
}

func TestMerge(t *testing.T) {
	arr1 := []int{9, 8, 3}
	arr2 := []int{6, 5}

	ans := merge(arr1, arr2)
	fmt.Printf("ans: %+v\n", ans)
}

func TestGreater(t *testing.T) {
	arr1 := []int{6, 5}
	arr2 := []int{9, 8, 3}

	want := greater(arr1, arr2, 0, 0)
	fmt.Printf("want: %t\n", want)
}

func TestMaxNumber(t *testing.T) {
	arr1 := []int{5, 5, 1}
	arr2 := []int{4, 0, 1}

	ans := maxNumber(arr1, arr2, 3)
	fmt.Printf("ans: %+v\n", ans)
}
