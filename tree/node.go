package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

/**
为结构定义方法
值接收
*/
func (node Node) Print() {
	fmt.Print(node.Value, "  ")
}

/**
为结构定义方法
指针接收
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func main() {
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)
	root.Print()
	fmt.Println()
	root.Traverse()
	fmt.Println()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()
	nodes := []Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	pRoot := &root
	pRoot.Print()
	fmt.Println()
	pRoot.SetValue(200)
	pRoot.Print()
	fmt.Println()
	root.Print()

	var nilPRoot *Node
	nilPRoot.SetValue(200)
	nilPRoot = &root
	nilPRoot.SetValue(300)
	nilPRoot.Print()
	root.Print()
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


包
1.每个目录一个包
2.main包包含可执行入口
3.为结构定义的方法必须放在同一个包内
4.可以是不同文件

如何扩充系统类型或者别人的类型
1.定义别名
2.使用组合
*/
