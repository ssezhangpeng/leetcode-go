package leetcode_0208

type Trie struct {
	Next  [26]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		Next:  [26]*Trie{},
		IsEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range []byte(word) {
		if node.Next[ch-'a'] == nil {
			node.Next[ch-'a'] = new(Trie)
		}
		node = node.Next[ch-'a']
	}
	node.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range []byte(word) {
		if node.Next[ch-'a'] == nil {
			return false
		}
		node = node.Next[ch-'a']
	}
	return node.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range []byte(prefix) {
		if node.Next[ch-'a'] == nil {
			return false
		}
		node = node.Next[ch-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
