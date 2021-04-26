package main

import (
	"fmt"
	"math"
	"sort"
)

func Sort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return int(math.Max(float64(nums[0]*nums[1]*nums[n-1]),
		float64(nums[n-1]*nums[n-2]*nums[n-3])))
}

func main() {
	fmt.Println(Sort([]int{1, 2, 3, 4, 5, 6}))
}
