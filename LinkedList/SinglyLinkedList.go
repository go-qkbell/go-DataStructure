package LinkedList

import "log"

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

func (s *SinglyLinkedList) AddAfter(n *SinglyNode, idx int) {
	if idx > s.length-1 {
		log.Println("cannot insert at this index!!")
		return
	}

	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		node := s.head
		for i := 0; i < idx; i++ {
			node = node.next
		}
		n.next = node.next
		node.next = n

		if node == s.tail {
			s.tail = n
		}
	}

	s.length++
}

func (s *SinglyLinkedList) RemoveFront() {
	if s.head == nil {
		return
	} else if s.head == s.tail {
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
	} else if s.head == s.tail {
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
