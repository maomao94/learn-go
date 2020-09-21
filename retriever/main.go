package main

import (
	"fmt"
	"learn-go/retriever/mock"
	"learn-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	} //指针接收者，必须使用地址接收
	inspect(r)
	fmt.Printf("%T %v\n", r, r) // 一个类型 一个值
	fmt.Println(download(r))

	// Type assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
