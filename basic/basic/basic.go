package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var () 定义
var (
	aa = 3
	ss = "kkk"
	bb = true
)

//bb := true 函数外面必须有关键字

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler1() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang) //0 4 2 3
	fmt.Println(b, kb, mb, gb, tb, pb)           //0 4 2 3
}

func basic() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	//验证欧拉公式
	euler()

	triangle()

	consts()

	enums()
}

/**
内建变量类型
1.bool string
2.(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
3.byte,rune
4.float32,float64,complex64,complex128


强制类型转换
1.类型转换是强制的


常量定义
const 数值可作为各种类型使用

常量定义枚举类型
1.普通枚举类型
2.自增值枚举类型



*/
