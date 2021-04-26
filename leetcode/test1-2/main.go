package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func recursion(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	new_head := recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return new_head
}

// 链表反转-递归
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
	node := recursion(&node1)
	fmt.Println("递归后", node)
}
