package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)
	result := Query(done)
	for v := range result {
		fmt.Println("query name ", v)
		if v == "宋飞" {
			return
		}
	}

	//done1 := make(chan struct{})
	//defer close(done1)
	//strings := a(done1)
	//for v := range strings {
	//	fmt.Println("main-", v)
	//}
}

func Query(done <-chan struct{}) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		tm := time.After(10 * time.Second)
		lists := []string{"张三", "李四", "小明", "宋飞", "王超", "马汉"}
		for {
			for _, v := range lists {
				select {
				case result <- v:
					fmt.Println("协程方法 query method", v)
				case <-done:
					fmt.Println("done")
					return
				case <-tm:
					fmt.Println("get name timeout")
					return
				}
			}
		}

	}()
	return result
}

func a(done <-chan struct{}) <-chan string {
	req := make(chan string)
	go func(<-chan string) {
		defer close(req)
		tm := time.After(1 * time.Second)
		tms := time.After(10 * time.Second)
		for {
			select {
			case <-tm:
				req <- "success"
				fmt.Println("a-success...")
			case <-done:
				fmt.Println("a-done...")
				return
			case <-tms:
				fmt.Println("a-tms-done...")
				return
			}
		}
	}(req)
	return req
}
