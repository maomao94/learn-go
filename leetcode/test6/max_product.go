package main

import (
	"fmt"
	"math"
	"sort"

	"golang.org/x/tools/container/intsets"
)

func Sort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return int(math.Max(float64(nums[0]*nums[1]*nums[n-1]),
		float64(nums[n-1]*nums[n-2]*nums[n-3])))
}

func getMaxMin(nums []int) int {
	min1, min2 := intsets.MaxInt, intsets.MaxInt
	max1, max2, max3 := intsets.MinInt, intsets.MinInt, intsets.MinInt
	for _, x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}
	return int(math.Max(float64(min1*min2*max1),
		float64(max1*max2*max3)))
}

func main() {
	fmt.Println(Sort([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(getMaxMin([]int{1, 2, 3, 4, 5, 6}))
}
