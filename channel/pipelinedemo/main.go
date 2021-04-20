package main

import (
	"fmt"
	"learn-go/channel/pipeline"
)

func main() {
	p := pipeline.ArraySource(3, 2, 6, 7, 4)
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
