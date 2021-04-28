package main

import "fmt"

func arrangeCoins(n int) int {
	for i := 1; i <= n; i++ {
		n = n - i
		if n <= i {
			return i
		}
	}
	return -1
}

// 排列硬币-三种解法
func main() {
	fmt.Println(arrangeCoins(10))
}
