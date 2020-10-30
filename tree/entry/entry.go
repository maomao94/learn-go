package entry

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

func (myTreeNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed")
}

func main() {
	//var root tree.Node
	//root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}} //修改 使用内嵌
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Print()
	fmt.Println("root.Traversal")
	root.Traverse()
	fmt.Println("root.Node.Traversal")
	root.Node.Traverse()
	fmt.Println()
	root.postOrder()
	//node := myTreeNode{&root}
	//node.postOrder()
	//fmt.Println()

	//父类指针不能指向子类
	//var baseRoot *tree.Node
	//baseRoot := &root

	//函数式变成
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
