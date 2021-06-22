package leetcode_0065

import (
	"log"
	"testing"
)

func TestIsNumber(t *testing.T) {
	s := "0e"
	ans := isNumber(s)
	want := false
	if ans == want {
		log.Println("Pass")
	} else {
		log.Fatal("fucking fucked...")
	}
}
