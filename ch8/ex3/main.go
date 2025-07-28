package main

import "fmt"

func main() {
	fmt.Println("======================")
	l := List[int]{}
	l.Add(1)
	l.Add(2)
	n := l.head
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
	fmt.Println("======================")
	l.Add(3)
	n = l.head
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
	fmt.Println("======================")
	l.Insert(4, 1)
	n = l.head
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
	fmt.Println("======================")
	l.Insert(20, 20)
	fmt.Println("======================")
	fmt.Println(l.Index(4))
	fmt.Println(l.Index(5))
	l.Add(5)
	fmt.Println(l.Index(5))
	fmt.Println("======================")
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

// Adds a new element to the end of the linked list.
func (l *List[T]) Add(v T) {
	if l.head == nil {
		n := Node[T]{val: v}
		l.head = &n
		l.tail = &n
		return
	}
	lastNode := l.tail
	lastNode.next = &Node[T]{val: v}
	l.tail = lastNode.next
}

// Adds a new element at the specified position in the linked list.
func (l *List[T]) Insert(v T, i int) {
	currentIndex := 0
	n := l.head
	for n != nil {
		if currentIndex+1 == i {
			n.next = &Node[T]{
				val:  v,
				next: n.next,
			}
			return
		}
		currentIndex++
		n = n.next
	}

	fmt.Println("Index is out of range.")
	return
}

// Returns the position of the supplied value, -1 if it's not present.
func (l *List[T]) Index(v T) int {
	currentIndex := 0
	n := l.head
	for n != nil {
		if n.val == v {
			return currentIndex
		}

		currentIndex++
		n = n.next
	}

	return -1
}
