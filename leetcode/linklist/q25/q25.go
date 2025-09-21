package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
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
