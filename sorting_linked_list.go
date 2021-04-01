package main

import "fmt"

type LinkNode struct {
	next *LinkNode
	prev *LinkNode
	data int
}

type LinkList struct {
	head  *LinkNode
	index int
}

//Add Method to Struct List
func (L *LinkList) InsertLink(val int) *LinkNode {
	list := &LinkNode{
		data: val,
	}
	L.index++
	L.head = list
	return list
}

func (l *LinkList) ShowBackwards() {
	node := l.head
	for node != nil {
		fmt.Printf("%+v\n", node.data)
		node = node.prev
	}

}

func (l *LinkList) ShowForwards() {
	node := l.head
	lastNode := l.head
	for node != nil {
		lastNode = node
		node = node.prev
	}
	node = lastNode
	for node != nil {
		fmt.Printf("%+v\n", node.data)
		node = node.next
	}

}

func (l *LinkList) sorting() {
	node := l.head
	lastNode := l.head
	for node != nil {
		lastNode = node
		node = node.prev
	}
	firstNode := lastNode
	for firstNode.prev != nil {

		firstNode = firstNode.next
	}

	current := firstNode
	index := firstNode.next
	for current != nil {
		for index != nil {
			if current.data > index.data {
				temp := current.data
				current.data = index.data
				index.data = temp
			}
			index = index.next

		}
		current = current.next
		index = firstNode
	}

	node = firstNode

	for node != nil {
		fmt.Printf("%+v\n", node.data)
		node = node.next
	}
}

func main() {
	link := LinkList{}

	node1 := link.InsertLink(0)

	node2 := link.InsertLink(10)
	node2.prev = node1
	node1.next = node2

	node3 := link.InsertLink(1)
	node3.prev = node2
	node2.next = node3

	node4 := link.InsertLink(5)
	node4.prev = node3
	node3.next = node4

	node5 := link.InsertLink(12)
	node5.prev = node4
	node4.next = node5

	node6 := link.InsertLink(6)
	node6.prev = node5
	node5.next = node6

	node7 := link.InsertLink(20)
	node7.prev = node6
	node6.next = node7

	node8 := link.InsertLink(8)
	node8.prev = node7
	node7.next = node8

	fmt.Println("---back to front--------")
	link.ShowBackwards()
	fmt.Println("-----front to back------")
	link.ShowForwards()
	fmt.Println("----Sorted-------")
	link.sorting()
}
