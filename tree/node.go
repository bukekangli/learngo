package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value} //返回了局部变量的地址
}

func (node Node) Print() { // 值接收者
	fmt.Printf("%d ", node.Value)
}

func (node *Node) SetValue(value int) { //指针接收者
	if node == nil {
		fmt.Println("Setting Value to nil node ignored")
		return
	}
	node.Value = value
}



