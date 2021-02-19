package linkedlist

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	currentNode := linkedList
	for currentNode != nil {
		nextNode := currentNode.Next
		for nextNode != nil && currentNode.Value == nextNode.Value {
			nextNode = nextNode.Next
		}

		currentNode.Next = nextNode
		currentNode = nextNode
	}
	return linkedList
}
