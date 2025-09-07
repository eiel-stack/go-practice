package main

import "fmt"

type MyCircularQueue struct {
	front, rear int
	elements    []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		elements: make([]int, k+1),
	}
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *MyCircularQueue) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear-1+len(q.elements))%len(q.elements)]
}

func main() {
	circularQueue := Constructor(3)
	fmt.Println(circularQueue.EnQueue(1)) // 预期 true
	fmt.Println(circularQueue.EnQueue(2)) // 预期 true
	fmt.Println(circularQueue.EnQueue(3)) // 预期 true
	fmt.Println(circularQueue.EnQueue(4)) // 预期 false
	fmt.Println(circularQueue.Rear())     // 预期 3
	fmt.Println(circularQueue.IsFull())   // 预期 true
	fmt.Println(circularQueue.DeQueue())  // 预期 true
	fmt.Println(circularQueue.EnQueue(4)) // 预期 true
	fmt.Println(circularQueue.Rear())     // 预期 4

}
