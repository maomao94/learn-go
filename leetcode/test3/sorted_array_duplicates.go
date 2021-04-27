package main

import "fmt"

// 双指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 删除排序数组中的重复项
func main() {
	fmt.Println(
		removeDuplicates([]int{0, 1}))
}
