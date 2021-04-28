package main

import (
	"fmt"
	"math"
)

// 二分查找
func binarySearch(x int) int {
	index, l, r := -1, 0, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			index = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}

// 牛顿迭代
func newton(x int) int {
	if x == 0 {
		return -1
	}
	return int(sqrt(float64(x), float64(x)))
}

func sqrt(i float64, x float64) float64 {
	res := (i + x/i) / 2
	if res == i || math.Abs(res-i) < 0.000000000000000000001 {
		return i
	} else {
		return sqrt(res, x)
	}
}

func main() {
	fmt.Println(binarySearch(24))
	fmt.Println(newton(24))
}
