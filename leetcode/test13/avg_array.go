package main

import (
	"fmt"
	"math"
)

// 可以使用双指针来做 本方法使用滑动窗口算法
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	len := len(nums)
	// 先统计第一个窗口的和
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < len; i++ {
		sum = sum - nums[i-k] + nums[i]
		max = int(math.Max(float64(sum), float64(max)))
	}
	return float64(max) / float64(k)
}

// 子数组最大平均数
func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
