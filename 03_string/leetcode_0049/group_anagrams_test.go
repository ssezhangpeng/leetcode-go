package leetcode_0049

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)

	fmt.Println("ans: ", ans)
}
