package main

import (
	"fmt"
	"learn/Go_course/interface/infra"
	"learn/Go_course/interface/testing"
)

//Go語言是個面向interface的語言，其interface極為靈活
func getRetriever() retriever{
	return infra.Retriever{}
}
func getTestingRetriever() retriever{
	return testing.Retriever{}
}

/*
?: something that can Get
我需要定義一個型別，適用不同的function，降低型別耦合度，interface就是很好的途徑。
定義一個retriever的interface，裡面有個Get方法，輸入string返回string
*/
type retriever interface{
	Get(string)string
}
func main() {
	//如果你想在一個slice中放int和放string，型別就命名為interface{}
	var a = []interface{}{}
	a = append(a,"aaa")
	a = append(a,1)
	
	var r retriever = getRetriever()
	var tr retriever = getTestingRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
	fmt.Println(tr.Get("https://www.imooc.com"))
}
