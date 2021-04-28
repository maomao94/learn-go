package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

// 深度优先
func minDepth(root *treeNode) int {
	if root == nil {
		return -1
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	min := math.MaxInt64
	if root.left != nil {
		min = int(math.Min(float64(minDepth(root.left)), float64(min)))
	}
	if root.right != nil {
		min = int(math.Min(float64(minDepth(root.right)), float64(min)))
	}
	return min + 1
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
	fmt.Println(minDepth(&node1))
}
