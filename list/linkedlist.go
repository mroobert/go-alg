package list

import "fmt"

func LinkListMain() {
	list := &LinkedList{}

	node1 := &Node{Value: 48}
	node2 := &Node{Value: 18}
	node3 := &Node{Value: 16}
	node4 := &Node{Value: 2}
	node5 := &Node{Value: 89}
	node6 := &Node{Value: 10}
	node7 := &Node{Value: 2}

	list.Prepend(node1)
	list.Prepend(node2)
	list.Prepend(node3)
	list.Prepend(node4)
	list.Prepend(node5)
	list.Prepend(node6)
	list.Prepend(node7)

	list.PrintListData()
	list.Delete(2)
	list.PrintListData()
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n

	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) Delete(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next != nil {
		if previousToDelete.Next.Value == value {
			previousToDelete.Next = previousToDelete.Next.Next
			l.Length--
		} else {
			previousToDelete = previousToDelete.Next
		}
	}
}

func (l *LinkedList) PrintListData() {
	toPrint := l.Head

	for i := 0; i < l.Length; i++ {
		fmt.Printf("%d ", toPrint.Value)
		toPrint = toPrint.Next
	}

	fmt.Println()
}
