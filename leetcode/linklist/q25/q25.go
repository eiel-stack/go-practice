package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		head, tail = groupReverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return dummy.Next
}

func groupReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head
	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return tail, head
}

func main() {
	// 示例1
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	k := 2
	head = reverseKGroup(head, k)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	// 输出: 2 1 4 3 5

	fmt.Println()

	// 示例2
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	k = 3
	head = reverseKGroup(head, k)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	// 输出: 3 2 1 4 5
}
