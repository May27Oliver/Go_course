package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int
//golang i/o提供input output的許多方法
//golang可以為函式型別增加reciver
func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0,io.EOF//表輸入檔的終點
	}
	s := fmt.Sprintf("%d\n", next)
	fmt.Println("s",s)
	a:=strings.NewReader(s)
	fmt.Println("strings.NewReader(s)",a)
	l,m:=a.Read(p)
	fmt.Println("strings.NewReader(s).Read(p)",l,m)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {//接口在這裡只是個型別，實作由intGen來完成
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fibonacci() intGen {
	a, b := 0, 1 //閉包
	return func() int {
		a, b = b, a+b
		return a
	}
}

//golang的函式也有接收器。
func main() {
	f := fibonacci()
	printFileContents(f)
}
