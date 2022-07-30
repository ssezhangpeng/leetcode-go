package trie

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func Constructor() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	for _, ch := range word {
		if t.Children[ch-'a'] == nil {
			t.Children[ch-'a'] = &Trie{}
		}
		t = t.Children[ch-'a']
	}
	t.IsEnd = true
}

func (t *Trie) search(word string) bool {
	for _, ch := range word {
		if t.Children[ch-'a'] == nil {
			return false
		}
		t = t.Children[ch-'a']
	}
	return t.IsEnd
}

func (t *Trie) startsWith(prefix string) bool {
	for _, ch := range prefix {
		if t.Children[ch-'a'] == nil {
			return false
		}
		t = t.Children[ch-'a']
	}
	return true
}
