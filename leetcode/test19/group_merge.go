package main

import "fmt"

// 深度优先
func getProvince(citys [][]int) int {
	return 0
}

// 省份数量
func main() {
	fmt.Println(getProvince([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // 2
	fmt.Println(getProvince([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // 3
}
