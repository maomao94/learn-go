package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func hasCycle(head *ListNode) bool {
	set := NewHashSet()
	for head.Next != nil {
		if !set.Add(head) {
			return true
		}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head.Next
	for slow != quick {
		if quick == nil || quick.Next == nil {
			return false
		}
		slow = slow.Next
		quick = quick.Next.Next
	}
	return true
}

// golang 没有set集合
type HashSet struct {
	set map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[interface{}]bool)}
}

func (set *HashSet) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *HashSet) Get(i interface{}) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *HashSet) Remove(i interface{}) {
	delete(set.set, i)
}

// 环形链表
func main() {
	var node5 = ListNode{5, nil}
	var node4 = ListNode{4, &node5}
	var node3 = ListNode{3, &node4}
	var node2 = ListNode{2, &node3}
	var node1 = ListNode{1, &node2}
	node5.Next = &node3
	fmt.Println(hasCycle(&node1))
	fmt.Println(hasCycle2(&node1))
}
