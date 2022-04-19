package linkedlist

import (
	"fmt"
)

type SinglyLinkedList struct {
	head   *SinglyNode
	tail   *SinglyNode
	length int
}

type SinglyNode struct {
	next  *SinglyNode
	value int
}

func NewSinglyList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func NewSinglyNode(v int) *SinglyNode {
	return &SinglyNode{
		next:  nil,
		value: v,
	}
}

func (s *SinglyLinkedList) AddFront(n *SinglyNode) {
	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		n.next = s.head
		s.head = n
	}

	s.length++
}

func (s *SinglyLinkedList) AddBack(n *SinglyNode) {
	if s.tail == nil {
		s.head = n
		s.tail = n
	} else {
		s.tail.next = n
		s.tail = n
	}

	s.length++
}

func (s *SinglyLinkedList) AddAt(n *SinglyNode, idx int) {
	if idx == 0 {
		s.AddFront(n)
		return
	} else if idx > s.length {
		s.AddBack(n)
		return
	} else {
		node := s.head
		for i := 0; i < idx-1; i++ {
			node = node.next
		}
		n.next = node.next
		node.next = n
	}

	s.length++
}

func (s *SinglyLinkedList) RemoveFront() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.head = s.head.next
	}

	s.length--
}

func (s *SinglyLinkedList) RemoveBack() {
	if s.tail == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		for n := s.head; n != nil; {
			if n.next == s.tail {
				s.tail = n
				n.next = nil
			}
			n = n.next
		}
	}

	s.length--
}

func (s *SinglyLinkedList) RemoveNode(n *SinglyNode) {
	if s.head == n {
		s.head = s.head.next
		s.length--
		return
	}

	for pn := s.head; pn != nil; {
		if pn.next == n {
			pn.next = n.next

			if n == s.tail {
				s.tail = pn
			}

			break
		}

		pn = pn.next
	}

	s.length--
}

func (s *SinglyLinkedList) Traverse() {
	for node := s.head; node != nil; {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
}
