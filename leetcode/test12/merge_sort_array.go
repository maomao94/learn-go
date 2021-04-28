package main

import (
	"fmt"
)

// 双指针  go语言copy方法的局限性
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	temp := make([]int, m)
	copy(temp, nums1)
	p1, p2 := 0, 0 //指向temp,nums2
	for i := 0; i < len(nums1); i++ {
		if p1 >= m {
			nums1[i] = nums2[p2]
			p2++
			continue
		}
		if p2 >= n {
			nums1[i] = temp[p1]
			p1++
			continue
		}
		if temp[p1] < nums2[p2] {
			nums1[i] = temp[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
	return nums1
}

// 合并两个有序数组
func main() {
	nums1 := []int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0}
	nums2 := []int{2, 4, 6, 8, 10}
	fmt.Println(merge(nums1, 5, nums2, 5))
}
