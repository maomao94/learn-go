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

// 二分查找
func arrangeCoins2(n int) int {
	low, high := 0, n
	for low <= high {
		mid := (high-low)/2 + low
		cost := ((mid + 1) * mid) / 2
		if cost == n {
			return mid
		} else if cost > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 排列硬币-三种解法
func main() {
	fmt.Println(arrangeCoins(10))
	fmt.Println(arrangeCoins2(10))
}
