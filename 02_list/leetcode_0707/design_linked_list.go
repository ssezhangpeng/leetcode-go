package leetcode0707

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type MyLinkedList struct {
	Length int
	Head   *ListNode
	Tail   *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Length {
		return -1
	}
	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Prev: nil,
		Next: this.Head,
	}
	if this.Length == 0 {
		this.Head = node
		this.Tail = node
	} else {
		this.Head.Prev = node
		this.Head = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Prev: this.Tail,
		Next: nil,
	}
	if this.Length == 0 {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.Next = node
		this.Tail = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if index == this.Length {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	// insert before p
	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	node := &ListNode{
		Val:  val,
		Prev: p.Prev,
		Next: p,
	}
	p.Prev.Next = node
	p.Prev = node
	this.Length++
}

func (this *MyLinkedList) DeleteAtHead() {
	newHead := this.Head.Next
	newHead.Prev = nil
	this.Head.Next = nil
	this.Head = newHead
	this.Length--
}

func (this *MyLinkedList) DeleteAtTail() {
	newTail := this.Tail.Prev
	this.Tail.Prev = nil
	newTail.Next = nil
	this.Tail = newTail
	this.Length--
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}

	if index == 0 {
		this.DeleteAtHead()
		return
	}

	if index == this.Length-1 {
		this.DeleteAtTail()
		return
	}

	// delete node p
	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	this.Length--
}

func (this *MyLinkedList) PrintMyLinkedList() {
	p := this.Head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
	fmt.Println("------------")
}
