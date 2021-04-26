package main

import (
	"fmt"
)

// 暴力算
func bf(n int) int {
	var count = 0
	for i := 2; i < n; i++ {
		if ok := isPrime(i); ok {
			count = count + 1
		}
	}
	return count
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 埃氏筛选
func eratosthenes(n int) int {
	// 建立一个标记位置 默认全是素数
	isPrime := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			count++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = true
			}
		}
	}
	return count
}

func main() {
	fmt.Println(bf(100))
	fmt.Println(eratosthenes(100))
}
