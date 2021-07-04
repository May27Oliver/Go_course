package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//bool,string
//(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
//byte 8位元,rune 32位元
//float32, float64, complex64, complex128

/*
按照Go語言規範，任何型別在未初始化時都對應一個零值：
布林型別是false，整型是0，字串是""，
而指標、函式、interface、slice、channel和map的零值都是nil。
*/

func euler() {
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

//類型轉換
func triangle() {
	var a, b int = 3, 4
	var c int
	//Sqrt return的是float64 不能附值給int，必須先經過強制轉換
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	//在go語言中，首字母大寫代表的是public的意思
	//const 可以作為各種類型被使用，舉例而言，這裡的a,b就不需要像上面經過float64()轉換
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚舉 enums
func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	//b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//euler()
	//triangle()
	//consts()
	enums()
}
