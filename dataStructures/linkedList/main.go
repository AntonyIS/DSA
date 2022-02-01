package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Adds a new node at the beginning of the node
func (l *LinkedList) Prepend(value int) {
	// Create the Node to be added into the linked list
	var newNode Node = Node{data: value}
	// Check if linked list has no node and add the new node
	if l.head == nil {

		// Add the new node into the linked list
		l.head = &newNode

	} else {

		// Linked list has some node in it, add new node
		newNode.next = l.head
		l.head = &newNode
	}
}

// Adds a new node at the end of the linked  list
func (l *LinkedList) Append(value int) {
	// Create the Node to be added into the linked list
	var newNode Node = Node{data: value}
	// Initialiaze the current node
	current := l.head

	// Loop through linked list until you get to the last node
	for current.next != nil {
		// Move from one node to the next
		current = current.next
	}
	// Get the last node
	current.next = &newNode
}

// Returns the size /number of nodes in LinkedList
func (l LinkedList) Size() int {
	// Check if linked list is empty
	if l.head == nil {
		// Empty linked list
		return 0
	}
	// Set current to set the initial size
	var count int

	// Initilialize current node
	current := l.head
	// Loop through the nodes and increament count
	for current.next != nil {
		count += 1
		current = current.next
	}
	count += 1
	// Return count/size
	return count
}

// Returns a bool true or false if LinkedList is empty or not
func (l LinkedList) IsEmpty() bool {
	return l.head == nil
}

// Searches for a node with a value and removes from the linked list
func (l *LinkedList) Remove(value int) {
	// Check of linked list is empty
	if empty := l.IsEmpty(); empty {
		fmt.Println("Linked list empty")
	} else {
		// Check if the node to remove is the head
		if l.head.data == value {
			// Remove the head of the Linked list
			current := l.head
			l.head = l.head.next
			current.next = nil
		} else {
			// Loop through linked list and remove the found value
			current := l.head
			// Track the previous node
			previousNode := current

			for current.next != nil {
				if current.data == value {
					// Change pointer for the previous node
					previousNode.next = current.next
					current.next = nil
					return
				} else {
					previousNode = current
					current = current.next
				}
			}
		}
	}
}

// Searches for a value , returuns bool
func (l LinkedList) Search(value int) bool {
	// Check of linked list is empty
	if empty := l.IsEmpty(); empty {
		fmt.Println("Linked list empty")
		return false
	} else {
		// Check if the node to remove is the head
		if l.head.data == value {
			// Remove the head of the Linked list
			current := l.head
			l.head = l.head.next
			current.next = nil
			return true
		} else {
			// Loop through linked list and remove the found value
			current := l.head
			// Track the previous node
			previousNode := current

			for current.next != nil {
				if current.data == value {
					// Change pointer for the previous node
					previousNode.next = current.next
					current.next = nil
					return true
				} else {
					previousNode = current
					current = current.next
				}
			}
			return false
		}
	}
}

// Inserts a node in a specfic index
func (l *LinkedList) Insert(value, index int) {
	// Check of linked list is empty
	if empty := l.IsEmpty(); empty {
		fmt.Println("Linked list empty")
	} else {
		// Check if the node to remove is the head
		if index == 0 {
			// Add the head of the Linked list
			l.Prepend(value)
			return
		}
		if index >= 1 {
			// Create the Node to be added into the linked list
			newNode := Node{data: value}
			// Loop through linked list and remove the found value
			current := l.head
			// Track the previous node
			previousNode := current
			position := index

			// Loop through linked list and compare index with position
			for position >= 1 {
				previousNode = current
				current = current.next
				position -= 1
			}

			previousNode.next = &newNode
			newNode.next = current

		}
	}
}

func (l LinkedList) PrintList() {
	// Check if linked list is empty
	if l.head == nil {
		fmt.Println([]int{})
	} else {
		// Create slice to store nodes for standard outpout
		var myLinkedList []int

		// Initialize current node
		current := l.head

		for current.next != nil {
			myLinkedList = append(myLinkedList, current.data)
			current = current.next
		}
		myLinkedList = append(myLinkedList, current.data)
		fmt.Println(myLinkedList)
	}

}

func main() {
	linkedList := LinkedList{}
	linkedList.Prepend(10)
	linkedList.Prepend(20)
	linkedList.Prepend(30)
	linkedList.Prepend(40)
	linkedList.Prepend(50)
	linkedList.Append(60)
	linkedList.Append(70)
	linkedList.Append(80)
	linkedList.Append(90)
	linkedList.Append(100)
	linkedList.PrintList()
	linkedList.Insert(22, 0)
	linkedList.PrintList()
}
