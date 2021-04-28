package main

import (
	"fmt"
	"math"
)

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
	low, high, mid := 0, n, -1
	for low <= high {
		mid = (high-low)/2 + low
		cost := ((mid + 1) * mid) / 2
		if cost == n {
			return mid
		} else if cost > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid
}

// 牛顿迭代 --精度会低
func arrangeCoins3(n int) int {
	if n == 0 {
		return -1
	}
	return int(sqrt(float64(n), float64(n)))
}

func sqrt(x float64, n float64) float64 {
	res := (x + (2*n-x)/x) / 2
	if res == x || math.Abs(res-x) < 0.000000000000000000001 {
		return x
	} else {
		return sqrt(res, n)
	}
}

// 排列硬币-三种解法
func main() {
	fmt.Println(arrangeCoins(10))
	fmt.Println(arrangeCoins2(10))
	fmt.Println(arrangeCoins3(10))
}
