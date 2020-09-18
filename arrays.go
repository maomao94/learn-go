package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历1
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//遍历2
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//遍历3
	for _, v := range arr3 {
		fmt.Println(v)
	}
}
