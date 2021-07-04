package main

import "fmt"

func changeValue(s *int) {
	*s = *s + 2
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	s := 4
	pts := &s
	changeValue(pts)
	fmt.Println(s) //6

	a, b := 3, 5
	swap(&a, &b)
	fmt.Println(a, b)
}
