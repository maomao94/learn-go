package main

import (
	"fmt"

	"github.com/chentaihan/container/queue"

	"github.com/chentaihan/container/stack"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
	deep  int
}

func main() {
	node7 := treeNode{value: 7, left: nil, right: nil}
	node6 := treeNode{value: 6, left: nil, right: nil}
	node5 := treeNode{value: 5, left: &node6, right: &node7}
	node4 := treeNode{value: 4, left: nil, right: nil}
	node3 := treeNode{value: 3, left: nil, right: nil}
	node2 := treeNode{value: 2, left: &node4, right: &node5}
	node1 := treeNode{value: 1, left: &node2, right: &node3}
	fmt.Print("前序遍历-")
	preorder(&node1)
	fmt.Print("\n前序遍历-")
	preorder2(&node1)
	fmt.Print("\n中序遍历-")
	midorder(&node1)
	fmt.Print("\n中序遍历-")
	midorder2(&node1)
	fmt.Print("\n后序遍历-")
	postorder(&node1)
	fmt.Print("\n后序遍历-")
	postorder2(&node1)
	fmt.Print("\n层序遍历-")
	//list := make([]int, 12)
	list := make([]int, 0)
	levelorder(&node1, 1, &list)
	fmt.Print(list)
	fmt.Print("\n层序遍历-")
	levelorder2(&node1)
}

// 前序遍历-递归
func preorder(root *treeNode) {
	if root == nil {
		return
	}
	// 第一次成为栈顶元素打印
	fmt.Print(root.value)
	preorder(root.left)
	preorder(root.right)
}

// 前序遍历-迭代
func preorder2(root *treeNode) {
	if root != nil {
		stack := stack.NewStackLink()
		stack.Push(root)
		for !stack.Empty() {
			a, _ := stack.Pop()
			if a.(*treeNode) != nil {
				fmt.Print(a.(*treeNode).value)
				stack.Push(a.(*treeNode).right)
				stack.Push(a.(*treeNode).left)
			}
		}
	}
}

// 中序遍历-递归
func midorder(root *treeNode) {
	if root == nil {
		return
	}
	midorder(root.left)
	fmt.Print(root.value)
	midorder(root.right)
}

// 中序遍历-迭代
func midorder2(root *treeNode) {
	if root != nil {
		stack := stack.NewStackLink()
		for !stack.Empty() || root != nil {
			if root != nil {
				stack.Push(root)
				root = root.left
			} else {
				a, _ := stack.Pop()
				root = a.(*treeNode)
				fmt.Print(root.value)
				root = root.right
			}
		}
	}
}

// 后序遍历-递归
func postorder(root *treeNode) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Print(root.value)
}

// 后序遍历-迭代 最麻烦的一种
func postorder2(root *treeNode) {
	if root != nil {
		stack := stack.NewStackLink()
		prev := new(treeNode)
		for !stack.Empty() || root != nil {
			for root != nil {
				stack.Push(root)
				root = root.left
			}
			a, _ := stack.Pop()
			root = a.(*treeNode)
			if root.right == nil || root.right == prev {
				fmt.Print(root.value)
				prev = root
				root = nil
			} else {
				stack.Push(root)
				root = root.right
			}
		}
	}
}

// 层序遍历-递归
func levelorder(root *treeNode, i int, list *[]int) {
	if root == nil {
		return
	}
	length := len(*list)
	if length <= i {
		for j := 0; j <= i-length; j++ {
			// 扩充容量
			*list = append(*list, -1)
		}
	}
	(*list)[i] = root.value
	levelorder(root.left, 2*i, list)
	levelorder(root.right, 2*i+1, list)
}

// 层序遍历-迭代
func levelorder2(root *treeNode) {
	q := queue.NewQueueLink()
	q.Enqueue(root)
	for !q.Empty() {
		a, _ := q.Dequeue()
		if a.(*treeNode) != nil {
			fmt.Print(a.(*treeNode).value)
			q.Enqueue(a.(*treeNode).left)
			q.Enqueue(a.(*treeNode).right)
		}
	}
}
