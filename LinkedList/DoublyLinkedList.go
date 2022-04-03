package main

import "log"

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

type DoublyNode struct {
	next  *DoublyNode
	prev  *DoublyNode
	value int
}

func NewDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func NewDoublyNode(v int) *DoublyNode {
	return &DoublyNode{
		next:  nil,
		prev:  nil,
		value: v,
	}
}

func (s *DoublyLinkedList) AddFront(n *DoublyNode) {
	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		n.next = s.head
		s.head.prev = n
		s.head = n
	}

	s.length++
}

func (s *DoublyLinkedList) AddBack(n *DoublyNode) {
	if s.tail == nil {
		s.head = n
		s.tail = n
	} else {
		n.prev = s.tail
		s.tail.next = n
		s.tail = n
	}

	s.length++
}

func (s *DoublyLinkedList) AddAt(n *DoublyNode, idx int) {
	if idx > s.length-1 {
		log.Println("cannot insert at this index!!")
		return
	}

	if idx == 0 {
		s.AddFront(n)
		return
	} else {
		node := s.head
		for i := 0; i < idx-1; i++ {
			node = node.next
		}
		n.next = node.next
		n.prev = node
		node.next.prev = n
		node.next = n

	}

	s.length++
}

func (s *DoublyLinkedList) RemoveFront() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.head = s.head.next
		s.head.prev = nil
	}

	s.length--
}

func (s *DoublyLinkedList) RemoveBack() {
	if s.tail == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.tail = s.tail.prev
		s.tail.next = nil
	}

	s.length--
}

func (s *DoublyLinkedList) RemoveNode(n *DoublyNode) {
	if s.head == n {
		s.RemoveFront()
		return
	}
	if n == s.tail {
		s.RemoveBack()
		return
	}

	pn := n.prev

	pn.next = pn.next.next
	pn.next.prev = pn

	s.length--
}

func (s *DoublyLinkedList) Search(idx int) *DoublyNode {
	if s.head == nil {
		return nil
	}
	if idx > s.length-1 {
		return nil
	}

	var node *DoublyNode

	for i := 0; i < idx; i++ {
		node = node.next
	}

	return node
}
