package main

import "fmt"

//数组是值类型  会拷贝数组
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//改成指针类型
func printArray1(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//func main() {
//	var arr1 [5]int
//	arr2 := [3]int{1, 3, 5}
//	arr3 := [...]int{2, 4, 6, 8, 10}
//	var grid [4][5]bool
//	fmt.Println(arr1, arr2, arr3)
//	fmt.Println(grid)
//
//	fmt.Println("printArray(arr3)")
//	//遍历1
//	for i := 0; i < len(arr3); i++ {
//		fmt.Println(arr3[i])
//	}
//
//	fmt.Println("printArray(arr3)")
//	//遍历2
//	for i, v := range arr3 {
//		fmt.Println(i, v)
//	}
//
//	fmt.Println("printArray(arr3)")
//	//遍历3
//	for _, v := range arr3 {
//		fmt.Println(v)
//	}
//
//	fmt.Println("printArray(arr1)")
//	printArray(arr1)
//	fmt.Println("printArray(arr3)")
//	printArray(arr3)
//	fmt.Println("printArray(arr1,arr3)")
//	fmt.Println(arr1, arr3)
//
//	fmt.Println("printArray(&arr1)")
//	printArray1(&arr1)
//	fmt.Println("printArray(&arr3)")
//	printArray1(&arr3)
//
//	fmt.Println("printArray(arr1,arr3)")
//	fmt.Println(arr1, arr3)
//}
