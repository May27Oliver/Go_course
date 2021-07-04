package main

import (
	"fmt"
)

func eval(a, b int, op string) int {
	//go switch不用加break，如果想自動執行下個條件要加fallthrough
	//go switch後面也可以不加表達式，每個case後面加上條件即可
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func goRange() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	const filename = "./controlFlow/abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil { //nil是空值
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	//golang if 可以如下寫，但如下寫，就不能在if外調用contents,err
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	goRange()
}
