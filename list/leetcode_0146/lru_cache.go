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

	fmt.Println("front: ", this.CacheList.Front().Value.(list.Element).Value.(*Node).Key)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.CacheMap[key]; !ok {
		if this.Capacity == this.Size {
			// replacement
			last := this.CacheList.Back()

			node, _ := last.Value.(list.Element).Value.(*Node)
			delete(this.CacheMap, node.Key)
			this.CacheList.Remove(last)
			this.Size--
		}
		elem := list.Element{Value: &Node{Key: key, Value: value}}
		this.CacheList.PushFront(elem)
		this.CacheMap[key] = &elem
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
