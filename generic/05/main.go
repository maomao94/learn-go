package main

import (
	"fmt"
	"strconv"
)

type Price int

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

type ShowPrice interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

func showPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// 使用interface约束泛型
// go run -gcflags=-G=3 main.go
func main() {
	fmt.Println(showPriceList([]Price{48, 88, 152, 219, 328}))
}
