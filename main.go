package main //所有的go程式都需要有package，說明這個文件在哪個package下

import (
	"fmt"
	"math"
	"reflect"
)

//函式外面只能寫變數，函式，物件類別的宣告，不能寫表達式
// var name string
// var age int
// var isOk  bool

//	大量聲明變數
var (
	name string // 對應類型空值 ""
	age  int    // 0
	isOk bool   // false
)

func foo(a int, b string) {
	return
}

//Go變數宣告方式，宣告全域變數用
var (
	student_name string
	studentName  string
	StudentName  string
)

//函式外面每個語句都要以關鍵字開始(var,const,func)
//:=不能使用在函式外
//_多用於佔位，表示忽略值

//常數，常數是恆常不變的變數
const pi = 3.14

//多筆聲明常數
const (
	statusOK = 200
	notFound = 404
)

//如果宣告多筆常數，沒有值的變數會默認跟上面的變數值一樣
const (
	n1 = 100
	n2
	n3
)

//iota iota只用在常數內，類似計數器的存在
const (
	a1 = iota //0
	a2
	a3
)

const (
	b1 = iota
	b2
	_ //iota會多計算一次，但不會存入變數
	b3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota
	c4
)

func main() { //程式進入點
	name = "國動"
	age = 35
	isOk = true
	//Go會提醒你變數要用，沒用到的變數會報錯
	fmt.Println(isOk)           //Print在終端機印出內容
	fmt.Printf("name:%s", name) //%s :站位符 使用name這個變量的值去替換站位符
	fmt.Println(age)            //Println是換行的意思

	//聲明變數同時賦值
	var s1 string = "abc"
	fmt.Print(s1)

	//類型推導,不用宣告類型依照賦值判斷類型
	var s2 = "20"
	fmt.Print(s2)

	//短變數宣告:=
	//只能在函式內用
	s3 := "巨鎚瑞斯"
	fmt.Print(s3)

	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b1", b3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	//int vs uint
	fmt.Printf("uint8  : 0 ~ %d\n", math.MaxUint8)
	fmt.Printf("uint16 : 0 ~ %d\n", math.MaxUint16)
	fmt.Printf("uint32 : 0 ~ %d\n", math.MaxUint32)
	fmt.Printf("uint64 : 0 ~ %d\n", uint64(math.MaxUint64))
	fmt.Printf("int8   : %d ~ %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  : %d ~ %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  : %d ~ %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  : %d ~ %d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("整數預設型態: %s\n", reflect.TypeOf(1))
}

//類別
//整數
//int8 int16 int32 int64
//uint8 uint16 uint32 uint64
//uint8 就是byte類型
