package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go LinkedList")
}

// Create a node : A node should have a data and a pointer to the next node if any else points to nil.
type Node struct {
	data      int   // Data that the node will contain
	next_node *Node // Pointer to the next node
}

// Create a LinkedList : A collection of nodes linked with pointers.
// A LinkedList must have a head of type Node to define the starting point.
type LinkedList struct {
	head *Node // The starting point of a linked list
}

// Check if a LinkedList is empty
func (ll LinkedList) is_empty() bool {
	// Checks if a LinkedList is empty or not
	// Returns a boolean true if empty else returns false
	return ll.head == nil // if LinkedList head is nill, the LinkedList is empty
}
