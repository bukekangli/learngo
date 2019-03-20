package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(100)
	s.Insert(10000)

	fmt.Println(s)
	fmt.Println(s.Has(1000))

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	//root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(root *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}
