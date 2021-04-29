package main

import "fmt"

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
	fmt.Print("\n中序遍历-")
	midorder(&node1)
	fmt.Print("\n后序遍历-")
	postorder(&node1)
}

// 前序遍历
func preorder(root *treeNode) {
	if root == nil {
		return
	}
	// 第一次成为栈顶元素打印
	fmt.Print(root.value)
	preorder(root.left)
	preorder(root.right)
}

// 中序遍历
func midorder(root *treeNode) {
	if root == nil {
		return
	}
	midorder(root.left)
	fmt.Print(root.value)
	midorder(root.right)
}

// 后序遍历
func postorder(root *treeNode) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Print(root.value)
}
