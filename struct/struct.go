package main

import (
	"fmt"
)

//struct宣告
type TreeNode struct {
	value       int
	left, right *TreeNode //在struct內要定義其他struct為型別，必須用pointer，否則會陷入無限迴圈
}

func (node TreeNode) print() { //要在struct新增method，必須在func後面添加接收器(node TreeNode)，稱之為Receiver，相當於js的this
	fmt.Print(node.value, " ")
}

func (node *TreeNode) setValue(s int) {
	node.value = s
}

func createNode(s int) *TreeNode {
	return &TreeNode{value: s}
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root TreeNode
	root = TreeNode{value: 3}
	root.left = &TreeNode{value: 4}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode) //&{0 <nil> <nil>}
	//new(TreeNode) === &TreeNode{}
	root.left.right = createNode(2)
	root.right.left.setValue(100)
	//root.right.left.print() //此處依舊可以印出來，因為go有語法糖將pointer類型自動轉換成(*p).print()

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(200)
	// pRoot.print()

	root.traverse()
}
