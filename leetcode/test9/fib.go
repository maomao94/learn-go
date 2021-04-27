package main

import "fmt"

// 暴力
func calculate1(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return calculate1(num-1) + calculate1(num-2)
}

// 去重递归
func calculate2(num int) int {
	arr := make([]int, num+1)
	return recurse(arr, num)
}

// 双指针迭代
func calculate3(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	low, high := 0, 1
	for i := 2; i <= num; i++ {
		sum := low + high
		low = high
		high = sum
	}
	return high
}

func recurse(arr []int, num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	if arr[num] != 0 {
		return arr[num]
	}
	arr[num] = recurse(arr, num-1) + recurse(arr, num-2)
	return arr[num]
}

func main() {
	fmt.Println(calculate1(10))
	fmt.Println(calculate2(10))
	fmt.Println(calculate3(10))
}
