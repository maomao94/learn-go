package main

import (
	"fmt"
	"strconv"
)

type Price int

type ShowPrice interface {
	String() string
}

func (i Price) String() string {
	return strconv.Itoa(int(i))
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
