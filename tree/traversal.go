package tree

func (node *Node) Traverse() { //todo: 遍历树结构
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
