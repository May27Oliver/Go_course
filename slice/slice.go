package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	//陣列slice，slice是基於array的映射
	// fmt.Println("arr[2:6]=", arr[2:6])
	// fmt.Println("arr[2:6]=", arr[:6])
	// fmt.Println("arr[2:6]=", arr[2:])
	// fmt.Println("arr[2:6]=", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	/*
		把陣列傳入func內會拷貝一個新的陣列在func作用域中執行，但slice不會，
		傳遞slice進func會修改基底的array
	*/
	updateSlice(s1)
	updateSlice(s2)
	//fmt.Println(s1) //[100 3 4 5 6]
	//fmt.Println(s2) //[100 1 100 3 4 5 6]

	//slice extends
	s3 := arr[2:6]
	s4 := s1[3:5] //請注意s1這裡本無6這個元素，但因為slice映射array的關係會印出6

	/*
		一個slice有三個值ptr,length,capacity
		ptr:該slice起始的元素
		length:該slice映射到幾個array元素
		capacity:該slice的ptr到映射的array最末位元素間有多少元素
	*/

	fmt.Println(s3)
	fmt.Println(s4)
}
