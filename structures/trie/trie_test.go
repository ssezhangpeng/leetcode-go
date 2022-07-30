package trie

import "testing"

func TestMain(t *testing.T) {
	obj := Constructor()

	s1 := "apple"
	obj.Insert(s1)
	if ans := obj.search(s1); ans != true {
		t.Errorf("Got %v expect %v", ans, true)
	}

	s2 := "app"
	if ans := obj.startsWith(s2); ans != true {
		t.Errorf("Got %v expect %v", ans, true)
	}

	t.Log("Test Finish...")
}
