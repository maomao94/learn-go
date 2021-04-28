package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
	deep  int
}

type QueueI []interface{}

func (q *QueueI) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *QueueI) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *QueueI) IsEmpty() bool {
	return len(*q) == 0
}

// 深度优先
func minDepth1(root *treeNode) int {
	if root == nil {
		return -1
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	min := math.MaxInt64
	if root.left != nil {
		min = int(math.Min(float64(minDepth1(root.left)), float64(min)))
	}
	if root.right != nil {
		min = int(math.Min(float64(minDepth1(root.right)), float64(min)))
	}
	return min + 1
}

// 广度优先
func minDepth2(root *treeNode) int {
	if root == nil {
		return -1
	}
	root.deep = 1
	queue := QueueI{}
	queue.Push(root)
	for !queue.IsEmpty() {
		node := queue.Pop().(*treeNode)
		if node.left == nil && node.right == nil {
			return node.deep
		}
		if node.left != nil {
			node.left.deep = node.deep + 1
			queue.Push(node.left)
		}
		if node.right != nil {
			node.right.deep = node.deep + 1
			queue.Push(node.right)
		}
	}
	return -1
}

// 二叉树最小深度
func main() {
	node7 := treeNode{value: 7, left: nil, right: nil}
	node6 := treeNode{value: 6, left: &node7, right: nil}
	node5 := treeNode{value: 5, left: nil, right: nil}
	node4 := treeNode{value: 4, left: nil, right: nil}
	node3 := treeNode{value: 3, left: &node6, right: nil}
	node2 := treeNode{value: 2, left: &node4, right: &node5}
	node1 := treeNode{value: 1, left: &node2, right: &node3}
	fmt.Println(minDepth1(&node1))
	fmt.Println(minDepth2(&node1))
}
