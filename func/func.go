package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//標準函式
func div(a, b int) int {
	return a + b
}

//回傳多個值
func divAndRest(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func switchCase(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int { //將函式當作參數傳遞給另一個參數
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func main() {
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
}
