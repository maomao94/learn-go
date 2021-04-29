package main

import "fmt"

func change(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else { // 20
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(change([]int{5, 5, 20}))
}
