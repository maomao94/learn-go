package main

import "fmt"

func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type Food struct {
	Name  string
	Price int
}

// 使用compareable内置类型比较
// go run -gcflags=-G=3 main.go
func main() {
	fmt.Println(index([]int{11, 22, 33, 44, 55}, 55))
	fmt.Println(index([]string{"红烧肉", "清蒸鱼", "大闸蟹", "九转大肠", "烤全羊"}, "九转大肠"))
	fmt.Println(
		index([]Food{
			{"红烧肉", 1},
			{"清蒸鱼", 2},
			{"大闸蟹", 3},
			{"九转大肠", 4},
			{"烤全羊", 5},
		}, Food{"清蒸鱼", 2}))
}
