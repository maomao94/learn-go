package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	//var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			//fmt.Println("Received from c1:", n)
			values = append(values, n)
		case n = <-c2:
			//fmt.Println("Received from c2:", n)
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}

func worker(id int, c chan int) {
	//for {
	//	//n := <-c
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d receiver %c\n", id, n)
	//}

	for n := range c {
		fmt.Printf("Worker %d receiver %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
