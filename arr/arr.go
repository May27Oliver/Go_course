package main

import (
	"fmt"
)

func changeArrValue(arr [5]int) {
	arr[0] = 100
}

func main() {
	// var arr1 [5]int
	// arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10} //...讓編譯器數有幾個數字
	// fmt.Println(arr1, arr2, arr3)
	changeArrValue(arr3)
	fmt.Println(arr3) //[2,4,6,8,10]，go語言內如果把arr丟進func裡，在func內會拷貝一份新的array，故在func裡面改變array任一單一值都不會影響外面的array
}
