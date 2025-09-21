package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {

	first := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	fourth := &ListNode{Val: 4}
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = second
	fmt.Println(hasCycle(first))
}
