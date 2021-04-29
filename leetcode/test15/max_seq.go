package main

import (
	"fmt"
	"math"
)

func findLength(nums []int) int {
	start := 0
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			start = i
		}
		max = int(math.Max(float64(max), float64(i-start+1)))
	}
	return max
}

// 最长连续递增序列-贪心算法
func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 3, 4, 3, 4, 5, 6, 7}))
}
