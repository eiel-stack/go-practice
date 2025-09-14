package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Previous node starts as nil
	curr := head       // Current node starts at the head

	// Travers the list
	for curr != nil {
		next := curr.Next // Save the next node
		curr.Next = prev  // Reverse the link

		// Move pointers forward
		prev = curr // Move prev to the current node
		curr = next // Move curr to the next node
	}
	// prev is now the new head of the reversed list
	return prev
}

// Added missing printList function
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

// Recursive Solution
func reverseListRecursive(head *ListNode) *ListNode {
	// Base case: empty list or single node
	if head == nil || head.Next == nil {
		return head
	}

	// Recursively reverse the rest of the list
	newHead := reverseListRecursive(head.Next)

	// Reverse the current node's pointer
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func main() {
	// Create a linked list
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}

	// Print the original list
	fmt.Println("Original List:")
	printList(head)

	// Reverse the list
	head = reverseList(head)

	// Print the reversed list
	fmt.Println("Reversed List:")
	printList(head)

	// Reverse the list
	head = reverseListRecursive(head)

	// Print the reversed list
	fmt.Println("Reversed List (Recursive):")
	printList(head)
}
