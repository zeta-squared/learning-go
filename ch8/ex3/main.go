package main

type Node[T comparable] struct {
	val T
	next *Node[T]
}

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}
