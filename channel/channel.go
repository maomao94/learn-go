package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		//n := <-c
		fmt.Printf("Worker %d receiver %c\n", id, <-c)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			//n := <-c
			fmt.Printf("Worker %d receiver %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
