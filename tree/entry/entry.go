package main

import (
	"fmt"
	"learn/Go_course/tree"
)

//組合式擴充
// type myTreeNode struct {
// 	node *tree.Node
// }

//embedded內嵌式擴充
type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil { //指myTreeNode的*tree.Node
		return
	}

	(&myTreeNode{myNode.Left}).postOrder()
	(&myTreeNode{myNode.Right}).postOrder()
	myNode.Print()
}

type Comic struct {
	Universe string
}

func (c Comic) ComicUniverse() string {
	return c.Universe
}

type Marvel struct {
	c Comic
}

type Cellphone struct {
	ID   int
	name string
}

type Iphone struct {
	Cellphone
}

type HTC struct {
	Cellphone
}

func main() {
	// var root tree.Node
	// root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{Value: 4}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //&{0 <nil> <nil>}
	//new(TreeNode) === &TreeNode{}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(100)
	//root.right.left.print() //此處依舊可以印出來，因為go有語法糖將pointer類型自動轉換成(*p).print()

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(200)
	// pRoot.print()

	// root.Traverse()
	root.postOrder()

	newIphone := Iphone{}
	fmt.Println("newIphone", newIphone)
}
