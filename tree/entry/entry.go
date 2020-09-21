package main

import (
	"fmt"
	"learn-go/tree"
)

//组合方式扩展
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	node := myTreeNode{myNode.node.Left}
	node.postOrder()
	treeNode := myTreeNode{myNode.node.Right}
	treeNode.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Print()
	fmt.Println()
	root.Traverse()
	fmt.Println()
	node := myTreeNode{&root}
	node.postOrder()
}
