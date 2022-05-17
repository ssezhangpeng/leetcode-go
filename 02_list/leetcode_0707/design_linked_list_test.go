package leetcode0707

import (
	"fmt"
	"testing"
)

func TestDesignLinkList(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.PrintMyLinkedList()

	linkedList.AddAtTail(3)
	linkedList.PrintMyLinkedList()

	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.PrintMyLinkedList()

	ans := linkedList.Get(1) //返回2
	fmt.Println("ans: ", ans)

	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.PrintMyLinkedList()

	ans = linkedList.Get(1) //返回3
	fmt.Println("ans: ", ans)
}
