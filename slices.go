package main

func updateSlice(s []int) {
	s[0] = 100
}

//func main() {
//	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//	fmt.Println("arr[2:6] = ", arr[2:6])
//	fmt.Println("arr[:6] = ", arr[:6])
//	s1 := arr[2:]
//	fmt.Println("arr[2:] = ", s1)
//	s2 := arr[:]
//	fmt.Println("arr[:] = ", s2)
//	fmt.Println("------------")
//	updateSlice(s1)
//	fmt.Println("arr[2:] = ", s1)
//	updateSlice(s2)
//	fmt.Println("arr[:] = ", s2)
//	fmt.Println("Reslice")
//	s2 = s2[:5]
//	fmt.Println(s2)
//	s2 = s2[2:]
//	fmt.Println(s2)
//	fmt.Println(arr)
//
//	/**
//	slice实现
//	1.ptr指向slice 开头的元素
//	2.len slice的长度 超过会越界
//	3.cap 从ptr开始到结束
//	4.slice 可以向后扩展，不能向前扩展
//	*/
//	fmt.Println("Extending slice")
//	arr[0], arr[2] = 0, 2
//	fmt.Println(arr)
//	s1 = arr[2:6]
//	s2 = s1[3:5]
//	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
//	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
//	fmt.Println(s1[3:6])
//
//	/**
//	slice 添加元素
//	添加元素超越cap，系统会重新分配更大的底层数组
//	由于值传递关系，必须接收append的返回值
//	*/
//	fmt.Println("------------")
//	s3 := append(s2, 10)
//	fmt.Println("s3", s3)
//	fmt.Println(arr)
//	s3[0] = 55
//	fmt.Printf("replace s3[0]=%d\n", s3[0])
//	fmt.Println("s3", s3)
//	fmt.Println(arr)
//	fmt.Println("------------")
//	s3[0] = 555
//	s4 := append(s3, 11) //分配一个新的数组
//	s5 := append(s4, 12) //分配一个新的数组
//	s6 := s5[2:5]
//	s3[0] = 5
//	fmt.Printf("replace s3[0]=%d\n", s3[0])
//	fmt.Println("s3,s4,s5,s6", s3, s4, s5, s6)
//	s6[0] = 1010 //新的数组
//	fmt.Println("s6", s6)
//	// s4 and s5 no longer view arr.
//	fmt.Println(arr)
//}
