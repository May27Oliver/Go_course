package main //所有的go程式都需要有package，說明這個文件在哪個package下

import (
	"fmt"
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
	student_name,
	studentName,
	StudentName string
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

//函式宣告
//(x,y int)外的型別為回傳值的型別
func add(x, y int) int {
	return x + y
}

//函式本身可以是另一個函式的參數
func handleDeal(fn func(float64, float64) float64) float64 {
	return fn(3, 5)
}

//函式可以回傳另一個函式
func fibonacci() func(x int) int {
	return func(x int) int {
		return 2
	}
}

//函式回傳函式也有閉包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//struct 可以定義物件，接著丟進參數內
type addOpts struct {
	x int
	y int
	z int
}

func addOption(opts addOpts) int {
	return opts.x + opts.y + opts.z
}

func paramPtr(pPtr *int) {
	*pPtr = *pPtr + 10
	fmt.Println(*pPtr)
}

//variable自動判斷型別
func VariableDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//variable自動判斷型別
func VariableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//func 回傳多個值
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
	var (
		x int64 = 10
		y int64 = 20 //如果這裡x和y型別int32 int64不一樣，數字相加會錯
	)
	fmt.Println(x + y)

	//複數
	//複數包涵兩種型態 complex64、complex128。
	var complexValue complex64
	complexValue = 1.2 + 12i
	complexValue2 := 1.2 + 12i
	complexValue3 := complex(3.2, 12)

	fmt.Println("complexValue =", complexValue)
	fmt.Println("complexValue2 =", complexValue2)
	fmt.Println("complexValue3 =", complexValue3)

	fmt.Println("complexValue3 實數 =", real(complexValue3))
	fmt.Println("complexValue3 虛數 =", imag(complexValue3))

	//函式內定義函式可以用表達式
	add := func(x, y int) int {
		return x + y
	}
	add(1, 2)

	//匿名函式
	func() {
		fmt.Println("Hello anonymous")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	addOption(addOpts{x: 1, y: 2}) //呼喚參數為struct的func時須在參數前加上type

	//回傳多筆資料
	multiReturn := func() (x int, y string, z int) {
		return 1, "Oliver", 2
	}
	box1, box2, box3 := multiReturn()
	fmt.Println(box1, box2, box3)

	//取得記憶體位置
	var w int = 3
	var wPtr *int = &w //指標資料型態，須加上*
	fmt.Println(wPtr)
	//反解指標變數
	fmt.Println(*wPtr) //在指標變數前加上*可以反解該指標變數

	var p int = 10
	paramPtr(&p)
}

//類別
//整數
//int8 int16 int32 int64
//uint8 uint16 uint32 uint64
//uint8 就是byte類型
