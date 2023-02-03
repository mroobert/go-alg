package search

import "fmt"

func BinarySearchMain() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(125)
	tree.Insert(175)
	tree.Insert(110)
	tree.Insert(130)
	tree.Insert(160)
	tree.Insert(180)
	tree.Insert(111)
	tree.Insert(1)

	fmt.Println("Search element found: ", tree.Search(180))
	fmt.Println("Search steps: ", count)
}

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	count++

	if n == nil {
		return false
	}
	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}
	return true
}
