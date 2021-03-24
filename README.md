# Go_course
## 線上寫Go網站
https://play.golang.org/

## Package
package main //所有的go程式都需要有package，說明這個文件在哪個package下

## Import 要引用哪個套件包
import "fmt"

## 函式外的區域
函式外面只能寫變數，函式，物件類別的宣告，不能寫表達式

```go
 var name string
 var age int
 var isOk  bool
```


##	大量聲明變數
```go
var (
	name string // 對應類型空值 ""
	age  int    // 0
	isOk bool   // false
)
```

func foo(a int, b string) {
	return
}

### Go變數宣告方式，宣告全域變數用

```go
var (
	student_name string
	studentName  string
	StudentName  string
)
```
    函式外面每個語句都要以關鍵字開始(var,const,func)
    :=不能使用在函式外 
    _多用於佔位，表示忽略值


## 常數，常數是恆常不變的變數
```go
const pi = 3.14
```
## 多筆聲明常數
```go
const (
	statusOK = 200
	notFound = 404
)
```

## 如果宣告多筆常數，沒有值的變數會默認跟上面的變數值一樣
```go
const (
	n1 = 100
	n2 //100
	n3 //100
)
```

## iota iota只用在常數內，類似計數器的存在
```go
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2        //1
	_ //iota會多計算一次，但不會存入變數
	b3        //3
)


const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
```
## main()
跟Java一樣，Go的程式進入點在main()函式內
```go
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
}
```

## 類別
### 整數
int8 int16 int32 int64

uint8 uint16 uint32 uint64

uint8 就是byte類型，亦稱無整數
```go
uint8  : 0 ~ 255
uint16 : 0 ~ 65535
uint32 : 0 ~ 4294967295
uint64 : 0 ~ 18446744073709551615
int8   : -128 ~ 127
int16  : -32768 ~ 32767
int32  : -2147483648 ~ 2147483647
int64  : -9223372036854775808 ~ 9223372036854775807
```

## Go強制性編譯風格規範
```go
package main
import "fmt"
func main()
{
  i:= 1
  fmt.Println("Hello World", i)
}

//---以上是錯誤寫法
//---以下是正確寫法

package main
import "fmt"
func main() {
  i:= 1
  fmt.Println("Hello World", i)
}
```
## Go有非強制性的編譯風格建議
可以輸入指令：

    go fmt 檔名.go
    go fmt

即可針對想整理編譯風格的檔案進行縮排整理，不給予檔名的話就是所有檔案都會進行整理

##  如果宣告一個變數不用，會報錯哦！

### 流程控制 For

Go只有一種迴「 for 」。

基本的for迴圈除了沒有()外，跟java js一樣，必須有{}

```go
package main 

import "fmt"

func main(){
  //1.
  for i: = 0; i < 10; i++ {
    sum += 1
  }
  fmt.Println(sum)
  //2.
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
  //3. while 寫法
  sum:=1
  for sum < 1000{
    sum += sum
  }
  fmt.Println(sum)

  //4. 無窮迴圈
  //如果省略了條件，迴圈便不會結束，因此表達無限迴圈方式如下
  for {

  }
}
```