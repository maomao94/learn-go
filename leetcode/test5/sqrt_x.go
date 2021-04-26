package main

import "fmt"

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
	index, l, r := -1, 0, x
	for {
		if l <= r {
			mid := l + (r-l)/2
			if mid*mid <= x {
				index = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		break
	}
	return index
}

func main() {
	fmt.Println(binarySearch(24))
}
