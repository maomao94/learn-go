package main

import (
	"fmt"
	"learn-go/retriever/mock"
	"learn-go/retriever/real"
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
	fmt.Println(download(r))
	r = real.Retriever{}
	fmt.Println(download(r))
}
