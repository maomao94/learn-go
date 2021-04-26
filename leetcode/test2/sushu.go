package main

import "fmt"

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

func main() {
	fmt.Println(bf(100))
}
