package main

import "fmt"

// 暴力
func calculate(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return calculate(num-1) + calculate(num-2)
}

func main() {
	fmt.Println(calculate(10))
}
