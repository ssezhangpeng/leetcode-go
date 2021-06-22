package leetcode_0038

import (
	"log"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	n := 2
	ans := countAndSay(n)
	want := "11"
	if ans != want {
		log.Printf("ans: %s, want: %s", ans, want)
		log.Fatal("fucking fucked")
	} else {
		log.Println("Success")
	}
}
