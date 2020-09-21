package main

import (
	"fmt"
	"learn-go/tree"
)

//组合方式扩展  Embedding 内嵌
type myTreeNode struct {
	//node *tree.Node
	*tree.Node //Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	//left := myTreeNode{myNode.node.Left}
	//left.postOrder()
	//right := myTreeNode{myNode.node.Right}
	//right.postOrder()
	//myNode.node.Print()

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
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
