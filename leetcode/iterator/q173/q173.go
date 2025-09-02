package main

import "fmt"

// TreeNode 定义二叉搜索树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator 定义二叉搜索树迭代器结构
type BSTIterator struct {
	stack []*TreeNode
	curr  *TreeNode
}

// Constructor 初始化 BSTIterator
func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: []*TreeNode{},
		curr:  root,
	}
}

// Next 返回下一个节点的值
func (this *BSTIterator) Next() int {
	for this.curr != nil {
		this.stack = append(this.stack, this.curr)
		this.curr = this.curr.Left
	}
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.curr = node.Right
	return node.Val
}

// HasNext 判断是否有下一个节点
func (this *BSTIterator) HasNext() bool {
	return this.curr != nil || len(this.stack) > 0
}

func main() {
	// 构建示例中的二叉搜索树
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}

	// 初始化迭代器
	iterator := Constructor(root)

	// 测试 Next 和 HasNext 方法
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
}
