package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go LinkedList")
	ll := LinkedList{}

	// Check if LinkedList is empty
	fmt.Println(ll.is_empty()) // true
	// Check the length of the LinkedList
	fmt.Println(ll.length()) // 0
	// Add nodes LinkedList
	ll.prepend(10)
	fmt.Println(ll.length()) // 1
	ll.prepend(9)
	fmt.Println(ll.length()) // 2
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

// Get the length of a LinkedList
func (ll LinkedList) length() int {
	// Returns the length(int) of items in a LinkedList
	// This a linear operation 0(n)
	count := 0 //Initial length of a LinkedList

	current := ll.head // Keep track of the current node in a LinkedList

	// Loop through LinkedList until we get to the tail node and return count
	for current != nil {
		// As long as the next_node is not nil
		// Increament count
		count += 1
		// Move current node to the next node
		current = current.next_node
	}

	return count
}

// Add a node at the beginning of a LinkedList
func (ll *LinkedList) prepend(data int) {
	// Create a node to add in the LinkedList
	new_node := Node{data: data}

	// Check if the head of the LinkedList is nil.
	// Add new node as head node of the LinkedList
	if ll.head == nil {
		ll.head = &new_node
	} else {
		// LinkedList has existing nodes, add node to the existing nodes in the LinkedList
		new_node.next_node = ll.head // Link new node to the head of the LinkedList
		ll.head = &new_node          // Define new head for the LinkedList
	}
}
