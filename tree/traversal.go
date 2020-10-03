package tree

import "fmt"

func (node *Node) Traverse() {
	//if node == nil {
	//	return
	//}
	//node.Left.Traverse()
	//node.Print()
	//node.Right.Traverse()

	//使用函数来遍历二叉树
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}
