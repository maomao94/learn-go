package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func iterate(headNode *Node) *Node {
	var prev, next, cur *Node
	// 当前标记头节点
	cur = headNode
	for {
		// 当标记节点的next为nil时，代表循环结束
		if cur == nil {
			break
		} else {
			// 将标记节点的next留存
			next = cur.Next
			// 将标记节点的下一个设置成前一个节点
			cur.Next = prev
			// 将标记节点设置成前一个节点
			prev = cur
			// 将标记节点向后移动一个
			cur = next
		}
	}
	return prev
}

// 链表反转-迭代
func main() {
	var node5 = Node{5, nil}
	var node4 = Node{4, &node5}
	var node3 = Node{3, &node4}
	var node2 = Node{2, &node3}
	var node1 = Node{1, &node2}
	var cur = &node1
	for {
		fmt.Println(cur.Value)
		if cur.Next == nil {
			break
		} else {
			cur = cur.Next
		}
	}
	node := iterate(&node1)
	fmt.Println("反转后", node)
}
