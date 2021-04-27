package main

import (
	"fmt"
)

func solution1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func solution2(nums []int, target int) []int {
	flag := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := flag[target-nums[i]]; ok {
			return []int{flag[target-nums[i]], i}
		} else {
			flag[nums[i]] = i
		}
	}
	return []int{}
}

// 两数之和
func main() {
	fmt.Println(solution1([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(solution2([]int{1, 2, 3, 4, 5, 6}, 10))
}
