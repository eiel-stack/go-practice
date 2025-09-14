package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1
	}
	return dummyHead.Next
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairsRecursive(newHead.Next)
	newHead.Next = head

	return newHead
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	newHead := swapPairs(head)
	for newHead != nil {
		print(newHead.Val, " ")
		newHead = newHead.Next
	}

	print("\n")
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	newHead = swapPairsRecursive(head)
	for newHead != nil {
		print(newHead.Val, " ")
		newHead = newHead.Next
	}
}
