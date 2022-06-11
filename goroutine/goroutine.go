package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"sync"
)

func main() {
	//var a [10]int
	//for i := 0; i < 100000; i++ {
	//	go func(i int) {
	//		for {
	//			//a[i]++
	//			//runtime.Gosched() //主动交出控制权
	//			fmt.Printf("Hello from goroutine %d\n", i)
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Minute)
	//fmt.Println(a)

	// mapreduce
	var lock sync.Mutex
	lock.Lock()
	var uids []int64
	mr.Finish(func() error {
		uids = append(uids, 1)
		return nil
	}, func() error {
		uids = append(uids, 2)
		return nil
	}, func() error {
		uids = append(uids, 3)
		return nil
	})
	lock.Unlock()
	fmt.Println(uids)
	fmt.Println("##########")
	var uidsR []int64
	v, _ := mr.MapReduce(func(source chan<- interface{}) {
		for i := range uids {
			source <- uids[i]
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		fmt.Println(item)
		writer.Write(item.(int64) + 1)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		for item := range pipe {
			uidsR = append(uidsR, item.(int64))
		}
		writer.Write(uidsR)
	})
	fmt.Println(v)
}
