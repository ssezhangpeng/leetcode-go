package leetcode_0146

import (
	"container/list"
	"fmt"
)

type Node struct {
	Key   int
	Value int
}

type LRUCache struct {
	CacheList *list.List
	CacheMap  map[int]*list.Element
	Capacity  int
	Size      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list.New(),
		make(map[int]*list.Element),
		capacity,
		0,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.CacheMap[key]; !ok {
		return -1
	}

	elem := this.CacheMap[key]

	this.CacheList.MoveToFront(elem)

	node, _ := elem.Value.(*Node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.CacheMap[key]; !ok {
		if this.Capacity == this.Size {
			// replacement
			lastElem := this.CacheList.Back()

			node, _ := lastElem.Value.(*Node)
			delete(this.CacheMap, node.Key)
			this.CacheList.Remove(lastElem)
			this.Size--
		}
		node := &Node{Key: key, Value: value}
		elem := this.CacheList.PushFront(node)
		this.CacheMap[key] = elem
		this.Size++
	} else {
		// update node's Value
		elem := this.CacheMap[key]
		node, _ := elem.Value.(*Node)
		node.Value = value
		this.CacheList.MoveToFront(elem)
	}
	fmt.Printf("%+v\n", this.CacheMap)
}
