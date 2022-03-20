package tree

import (
	"fmt"
)

//struct宣告
type Node struct {
	Value       int
	Left, Right *Node //在struct內要定義其他struct為型別，必須用pointer，否則會陷入無限迴圈
}

func (node Node) Print() { //要在struct新增method，必須在func後面添加接收器(node TreeNode)，稱之為Receiver，相當於js的this
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(s int) {
	node.Value = s
}

func CreateNode(s int) *Node {
	return &Node{Value: s}
}
