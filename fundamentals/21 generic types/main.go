package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
}

func printLinkedList[T any](head *Node[T]) {
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func numbers() {
	tail := Node[int]{11, nil}
	node4 := Node[int]{7, &tail}
	node3 := Node[int]{5, &node4}
	node2 := Node[int]{3, &node3}
	head := Node[int]{2, &node2}
	printLinkedList(&head)
}

func letters() {
	tail := Node[string]{"West", nil}
	node3 := Node[string]{"South", &tail}
	node2 := Node[string]{"East", &node3}
	head := Node[string]{"North", &node2}
	printLinkedList(&head)
}

func main() {
	numbers()
	letters()
}
