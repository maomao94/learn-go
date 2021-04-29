package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(a []int) int {
	sort.Ints(a)
	for i := len(a) - 1; i >= 2; i-- {
		if a[i-1]+a[i-2] > a[i] {
			return a[i-1] + a[i-2] + a[i]
		}
	}
	return 0
}

// 三角形的最大周长-贪心算法
func main() {
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
}
