package main

import "fmt"

func main(){
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	q.Push(5)
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Pop()
	fmt.Print(q.IsEmpty())
}

type Queue []int

func (q *Queue) Push(v int){
	*q = append(*q,v)
}

func (q *Queue) Pop() int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool{
	return len(*q)== 0  
}