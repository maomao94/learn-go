package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

/**
为结构定义方法
值接收
*/
func (node treeNode) print() {
	fmt.Print(node.value, "  ")
}

/**
为结构定义方法
指针接收
*/
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.print()
	fmt.Println()
	root.traverse()
	fmt.Println()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	pRoot := &root
	pRoot.print()
	fmt.Println()
	pRoot.setValue(200)
	pRoot.print()
	fmt.Println()
	root.print()

	var nilPRoot *treeNode
	nilPRoot.setValue(200)
	nilPRoot = &root
	nilPRoot.setValue(300)
	nilPRoot.print()
	root.print()
}

/**
面向对象
1.go语言仅支持封装，不支持继承和多态
2.go语言没有class，只有struct

结构的创建
1.不论地址还是结构本身，一律使用 . 来访问成员
2.使用自定义工厂函数
3.注意返回了局部变量的地址
4.不需要知道结构创建在堆还是栈

为结构定义方法
1.显示定义和命名方法接收者
2.只有使用指针才可以改变结构的内容
3.nil 指针也可以调用方法

值接收者 vs 指针接收者
1.要改变内容必须使用指针接收者
2.结构过大也考虑指针接收者
3.一致性，如有指针接收者，最好都是指针接收者

值接受者 go语言特有
*/
