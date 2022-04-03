package main

import "testing"

func TestNewSinglyList(t *testing.T) {
	list := NewSinglyList()

	n0 := NewSinglyNode(0)
	n1 := NewSinglyNode(1)
	n2 := NewSinglyNode(2)
	n3 := NewSinglyNode(3)
	n4 := NewSinglyNode(4)
	n5 := NewSinglyNode(5)
	n6 := NewSinglyNode(6)
	n7 := NewSinglyNode(7)
	n8 := NewSinglyNode(8)
	n9 := NewSinglyNode(9)

	list.AddBack(n1)
	list.AddFront(n6)
	list.AddBack(n2)
	list.AddBack(n3)
	list.AddFront(n7)
	list.AddBack(n4)
	list.AddAt(n8, 2)
	list.AddAt(n9, 10)
	list.AddBack(n5)
	list.AddFront(n0)

	guess := []int{0, 7, 6, 8, 1, 2, 3, 4, 9, 5}
	count := 0

	for node := list.head; node != nil; {
		if node.value != guess[count] {
			t.Fatalf("node: %d, expected: %d failed", node.value, guess[count])
		}

		t.Logf("node: %d, expected: %d success", node.value, guess[count])
		node = node.next
		count++
	}
}

func TestNewDoublyList(t *testing.T) {
	list := NewDoublyList()

	n0 := NewDoublyNode(0)
	n1 := NewDoublyNode(1)
	n2 := NewDoublyNode(2)
	n3 := NewDoublyNode(3)
	n4 := NewDoublyNode(4)
	n5 := NewDoublyNode(5)
	n6 := NewDoublyNode(6)
	n7 := NewDoublyNode(7)
	n8 := NewDoublyNode(8)
	n9 := NewDoublyNode(9)

	list.AddBack(n1)
	list.AddFront(n6)
	list.AddBack(n2)
	list.AddBack(n3)
	list.AddFront(n7)
	list.AddBack(n4)
	list.AddAt(n8, 2)
	list.AddAt(n9, 10)
	list.AddBack(n5)
	list.AddFront(n0)

	guess := []int{0, 7, 6, 8, 1, 2, 3, 4, 9, 5}
	count := 0

	for node := list.head; node != nil; {
		if node.value != guess[count] {
			t.Fatal("failed")
		}

		t.Logf("node: %d, expected: %d", node.value, guess[count])
		node = node.next
		count++
	}
}
