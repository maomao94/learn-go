package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)
	fmt.Println("------------")
	updateSlice(s1)
	fmt.Println("arr[2:] = ", s1)
	updateSlice(s2)
	fmt.Println("arr[:] = ", s2)
}
