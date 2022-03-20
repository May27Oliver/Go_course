package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Retriever struct {}

func (Retriever) Get(url string)string{
	//如何讓系統可配置可測試進行併發
	//interface
	resp,err:=http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()//最後要close

	bytes,_:=ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n",bytes)
	return string(bytes)

}