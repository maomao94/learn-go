package main

import "fmt"

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

// 泛型基本语法
// go run -gcflags=-G=3 main.go
func main() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[float64]([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	printSlice[string]([]string{"红烧肉", "清蒸鱼", "大闸蟹", "九转大肠", "重烧海参"})
	printSlice([]int64{55, 44, 33, 22, 11})
}