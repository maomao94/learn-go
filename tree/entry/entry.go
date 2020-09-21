package main

import (
	"fmt"
	"learn-go/tree"
)

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
}
