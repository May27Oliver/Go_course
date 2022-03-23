package main

import (
	"fmt"
	"learn/Go_course/retriever/mock"
	"learn/Go_course/retriever/real"
	"time"
)

/*
1.interface實例化裡面有type也有值，一般來說在實例化interface可以不用用指針，因為interface自帶指針。
你使用值傳遞的interface，實例時可以用指針，但你在定義interface時若是以指針定義，實例時就只能是指針。
2.在golang裡面interface{}相當於typescript的any類型。
3.interface可以使用type assertion，type assertion型別斷言會回傳該interface的斷言變數，和一個斷言是否正確的ok值，若斷言變數與實際不符，ok值就會是false。
4.type switch可以
*/
type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster){
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":"Oliver",
			"learn":"Golang",
		},
	)
}

type RetrieverPoster interface{//interface組合，在interface內可以包含其他interface
	Retriever
	Poster
}
const url = "http://www.imooc.com"
func session(s RetrieverPoster)string{
	s.Post(url,map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is fake imooc.com"}
	r = &retriever
	inspect(r)
	//fmt.Printf("%T %v\n",r,r)//T：型別，v值
	r = &real.Retriever{//interface實例化後有兩個東西，一個是類型，一個是值
		UserAgent: "Mozilla/5.0",
		Timeout:time.Minute,
	}
	//fmt.Printf("%T %v\n",r,r)
	inspect(r)
	//type assertion型別斷言，在interface實例後加上.(型別)，型別斷言會回傳一個interface值，和一個檢驗是否是該型別的ok值。
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)
	if mockRetriever,ok := r.(*mock.Retriever); ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	//mock.retriever內有實作Post和Get方法，所以可以丟到Session內，刪掉任一個都會讓上面的程式碼出錯。
}

func inspect(r Retriever){
	switch v := r.(type){//interface的型別
		case *mock.Retriever:
			fmt.Println("Contents:",v.Contents)
		case *real.Retriever:
			fmt.Println("UseAgent:",v.UserAgent)
	}
}