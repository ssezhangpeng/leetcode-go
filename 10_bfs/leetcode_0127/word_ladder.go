package leetcode_0127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]bool)
	for _, word := range wordList {
		set[word] = true
	}

	if _, ok := set[endWord]; !ok {
		return 0
	}

	var queue []string
	queue = append(queue, beginWord)

	step := 0
	for len(queue) > 0 {
		step++
		size := len(queue) // 遍历每一层的元素
		for i := 0; i < size; i++ {
			word := queue[0]
			queue = queue[1:]

			for j := 0; j < len(word); j++ { // 尝试修改 word 中的每一个元素
				wordSlice := []byte(word)
				ch := wordSlice[j]
				// 暂存修改
				for k := 'a'; k <= 'z'; k++ {
					wordSlice[j] = byte(k)
					newWord := string(wordSlice)
					if newWord == endWord {
						return step + 1
					}
					if _, ok := set[newWord]; !ok {
						continue
					}

					delete(set, newWord)
					queue = append(queue, newWord)
				}
				// 恢复到修改之前的状态
				wordSlice[j] = ch
				word = string(wordSlice)
			}
		}
	}
	return 0
}
