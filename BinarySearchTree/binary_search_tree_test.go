package BinarySearchTree

import (
	"testing"
)

func TestNode_Insert(t *testing.T) {
	node := Node{5, nil, nil}
	node.Insert(10)
	node.Insert(1)
	node.Insert(4)
	node.Insert(100)
	node.Insert(20)
	node.Insert(0)
	node.Insert(2)
	node.Insert(9)

	list := make([]int, 0)
	l := node.InOrder(&list)
	t.Log(*l)
}
