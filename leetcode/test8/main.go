package main

import (
	"fmt"
)

// 二分查找
func twoSearch(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		low, high := i, len(nums)-1
		for low <= high {
			mid := (high-low)/2 + low
			if nums[mid] == target-nums[i] {
				return []int{i, mid}
			} else if nums[mid] > target-nums[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{}
}

// 双指针
func twoPoint(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum == target {
			return []int{low, high}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}

	return []int{}
}

// 两数之和-有序数组
func main() {
	fmt.Println(twoSearch([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(twoPoint([]int{1, 2, 3, 4, 5, 6}, 10))
}
