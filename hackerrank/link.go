// Linked List in Golang
package main

import "fmt"

type Node struct { // define a node
	prev *Node
	next *Node
	key  interface{}
}

type List struct { //  define a linked list
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {

	node := &Node{ // create a node
		next: L.head,
		key:  key,
	}

	if L.head != nil { // if linked list is not empty
		L.head.prev = node
	}
	L.head = node // set new head for linked list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l // set new tail for linked list
}

func (l *List) Display() {
	headNode := l.head    // address of 1st node in the linked list
	for headNode != nil { // iterate until  address is nil
		fmt.Printf("%+v ->", headNode.key) //print value stored in current node
		headNode = headNode.next           // move to next node
	}
	fmt.Println()
}

func Display(head *Node) {
	for head != nil {
		fmt.Printf("%v ->", head.key)
		head = head.next
	}
	fmt.Println()
}

func ShowBackwards(tail *Node) {
	for tail != nil {
		fmt.Printf("%v <-", tail.key)
		tail = tail.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert("a")
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	ShowBackwards(link.tail)
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	fmt.Println("\n==============================\n")
}
