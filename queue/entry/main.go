package main

import (
	"fmt"
	"learn-go/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println("---------")
	qi := queue.QueueI{"a"}
	qi.Push("abd")
	qi.Push("abdc")
	qi.Push1("abdce")
	fmt.Println(qi.Pop())
	fmt.Println(qi.Pop())
	fmt.Println(qi.IsEmpty())
	fmt.Println(qi.Pop())
	fmt.Println(qi.IsEmpty())
	fmt.Println(qi.Pop())
	fmt.Println(qi.IsEmpty())
}
