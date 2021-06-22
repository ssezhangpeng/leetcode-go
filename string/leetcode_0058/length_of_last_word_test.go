package leetcode_0058

import (
	"log"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	s := "a "

	ans := lengthOfLastWord(s)
	want := 1

	if ans != want {
		log.Printf("ans: %d, want: %d", ans, want)
		log.Fatal("fucking fucked")
	} else {
		log.Println("Success")
	}
}
