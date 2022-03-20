package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap = %d\n", len(s), cap(s))
}

func main() {
	var s []int //定義slice，Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s) //此處len會隨著回圈數遞增，但cap則會在與len相同數量後以倍數成長
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//陣列slice，slice是基於array的映射
	// fmt.Println("arr[2:6]=", arr[2:6])
	// fmt.Println("arr[2:6]=", arr[:6])
	// fmt.Println("arr[2:6]=", arr[2:])
	// fmt.Println("arr[2:6]=", arr[:])

	// s1 := arr[2:]
	// s2 := arr[:]
	/*
		把陣列傳入func內會拷貝一個新的陣列在func作用域中執行，但slice不會，
		傳遞slice進func會修改基底的array
	*/
	// updateSlice(s1)
	// updateSlice(s2)
	//fmt.Println(s1) //[100 3 4 5 6]
	//fmt.Println(s2) //[100 1 100 3 4 5 6]

	//slice extends
	// s3 := arr[2:6]
	// s4 := s1[3:5] //請注意s1這裡本無6這個元素，但因為slice內capacity會映射原陣列，故仍舊印出，但s1[i]不可超過len(s1)

	/*
		一個slice有三個值ptr,length,capacity
		ptr:該slice起始的元素
		length:該slice映射到幾個array元素
		capacity:該slice的ptr到映射的array最末位元素間有多少元素
	*/

	// fmt.Printf("s3 = %v ,len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))
	// fmt.Printf("s4 = %v ,len(s4) = %d ,cap(s3) = %d\n", s4, len(s4), cap(s4))

	// s5 := append(s4, 10)
	// s6 := append(s5, 11)
	// s7 := append(s6, 12)
	// fmt.Printf("s5 = %v ,len(s5) = %d, cap(s5) = %d\n", s5, len(s5), cap(s5))
	// fmt.Printf("s6 = %v ,len(s6) = %d, cap(s6) = %d\n", s6, len(s6), cap(s6))
	// fmt.Printf("s7 = %v ,len(s7) = %d, cap(s7) = %d\n", s7, len(s7), cap(s7))
	/*
		s5,s6,s7 view different array from origin array
		append時如果元素超越cap，系統會重新分配一個新的array給slice
		append必須有個新的變數接收
		s := append(s,val)
	*/
	fmt.Println("Deleting elements from slice")
	s8 := []int{2, 4, 6, 8, 10, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	s8 = append(s8[:3], s8[4:]...)
	fmt.Println(s8) //[2 4 6 10 12 0 0 0 0 0 0 0 0 0 0 0 0]

	//slice刪除第一個元素
	fmt.Println("Poping from front")
	s8 = s8[1:]
	//slice刪除最末位元素
	fmt.Println("Poping from the end")
	s8 = s8[:len(s8)-1]

}
