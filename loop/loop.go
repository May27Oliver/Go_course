package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

	/*
		strconv 將簡單的數據轉為字串型態
		strconv.Itoa() int轉字串
		strconv.Atoi() 字串轉int
	*/
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	//go 沒有while語法，直接用for 不加起始與遞增，就是意思等同while的語句了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//標準範例
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	//課程範例
	// fmt.Println(
	// 	convertToBin(5),
	// 	convertToBin(13),
	// 	convertToBin(72387885),
	// 	convertToBin(0),
	// )
	printFile("./loop/abc.txt")
}
