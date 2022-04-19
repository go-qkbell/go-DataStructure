package BinarySearchTree

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(key int) {
	if n.key < key {
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else if n.key > key {
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.key < key {
		return n.right.Search(key)
	} else if n.key > key {
		return n.left.Search(key)
	}

	return true
}
