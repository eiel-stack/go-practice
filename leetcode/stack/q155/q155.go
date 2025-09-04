package main

import (
	"fmt"
	"math"
)

// MinStack 定义最小栈结构
type MinStack struct {
	stack    []int // 主栈，用于存储所有元素
	minStack []int // 辅助栈，用于存储对应主栈位置的最小值
}

// Constructor 初始化最小栈
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt32}, // 辅助栈初始放入一个极大值，方便后续比较
	}
}

// Push 元素入栈
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 辅助栈添加当前的最小值（主栈当前所有元素的最小值）
	minVal := min(val, this.minStack[len(this.minStack)-1])
	this.minStack = append(this.minStack, minVal)
}

// Pop 元素出栈
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

// Top 获取栈顶元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin 获取栈中最小元素
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// min 辅助函数，返回两个整数中的较小者
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // 返回 -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // 返回 0
	fmt.Println(minStack.GetMin()) // 返回 -2
}
