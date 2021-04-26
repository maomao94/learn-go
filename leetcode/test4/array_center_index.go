package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		if total == sum {
			return i
		}
		sum = sum - nums[i]
	}
	return -1
}

func pivotIndex2(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		total := 2*left + nums[i]
		if total == sum {
			return i
		} else {
			left = left + nums[i]
		}
	}
	return -1
}

// 寻找数组的中心下表
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex2([]int{1, 7, 3, 6, 5, 6}))
}
