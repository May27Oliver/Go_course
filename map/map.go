package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	m := map[string]string{
		"name":    "oliver",
		"course":  "golang",
		"site":    "udemy",
		"quality": "good",
	}
	fmt.Println(m)

	m2 := make(map[string]int) //m2 ==empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m2, m3)

	//遍歷map
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//取得value
	courseName, ok := m["course"]                         //第二個變數ok接收true/false表這個key值是否存在map內
	fmt.Println("courseName = ", courseName, " ok =", ok) //courseName =  golang  ok = true

	//刪除value
	delete(m, "name")
	name, ok := m["name"]
	fmt.Println("name =", name, "ok = ", ok) //name =  ok =  false

	s := "我愛慕課網" //UTF-8
	for _, v := range []rune(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	fmt.Println("utf8.RuneCountInString(s)", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}

	//使用utf8.RuneCountInString來計算字串長度，len()只會計算字節數
	//使用[]byte取得字節
}
