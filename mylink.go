package main

import "fmt"

type Node struct {
	curr *Node
	data string
	prev *Node
	next *Node
}

type List struct {
	head  *Node
	index int
}

//Add Method to Struct List
func (L *List) Insert(name string, prevN *Node) *Node {
	list := &Node{
		curr: L.head,
		data: name,
		prev: prevN,
	}
	L.index++
	L.head = list
	return list
}

//Add Method to Struct List
func (l *List) ShowForwards() {
	list := l.head
	for list.prev != nil {
		list = list.prev
	}
	for list.next != nil {
		fmt.Printf("<-%+v\n", list)
		list = list.next
	}

}
func (l *List) ShowBackwards() {
	node := l.head
	for node.curr != nil {
		fmt.Printf("<-%+v\n", node)
		node = node.curr
	}
}

func main() {
	link := List{}
	node0 := link.Insert("", nil)
	node1 := link.Insert("Denmark", node0.curr)
	node2 := link.Insert("India", node1.curr)
	node1.next = node2
	node3 := link.Insert("UK", node2.curr)
	node2.next = node3
	node4 := link.Insert("USA", node3.curr)
	node3.next = node4
	node5 := link.Insert("Aus", node4.curr)
	node4.next = node5
	node6 := link.Insert("France", node5.curr)
	node5.next = node6
	node7 := link.Insert("Germany", node6.curr)
	node6.next = node7

	fmt.Println("\n==========reverse order================")
	link.ShowBackwards()

	fmt.Println("\n==========forward order================")
	link.ShowForwards()

	fmt.Println()
}
