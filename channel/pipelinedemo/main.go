package main

import (
	"bufio"
	"fmt"
	"learn-go/channel/pipeline"
	"os"
)

func main() {
	mergeDemo()

	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	write := bufio.NewWriter(file)
	pipeline.WriteSink(write, p)
	write.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	for v := range p {
		fmt.Println(v)
	}
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}

	// 写法2
	for v := range p {
		fmt.Println(v)
	}
}
