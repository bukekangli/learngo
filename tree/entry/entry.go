package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"learngo/tree"
	"time"
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
	fmt.Println(time.Now().Format("2006-01-02"))
	sql := "slk hahah %d"
	fmt.Printf(sql, 1)

	items := map[int64]string{
		1: "s",
		2: "l",
		3: "k",
		10: "shang",
		11: "SHANG",
		21: "li",
		22: "LI",
	}
	titleCmdSort := []int64{1, 2, 3, 10, 11, 21, 22}
	titleCmdMutex := make(map[int64]int64, len(items))
	l := make([]string, 0, len(items))
	for _, sort := range titleCmdSort {
		if _, ok := titleCmdMutex[sort/10]; !ok{
			l = append(l, items[sort])
			titleCmdMutex[sort/10] = 1
		}
	}
	fmt.Println()
	fmt.Println(l)
	fmt.Println()
	today, _ := time.Parse("2006-01-02", "2019-04-25")
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(today.Format("2006-01-02"))
	weekday := int(today.Weekday())
	m, _ := time.ParseDuration(fmt.Sprintf("-%dh", (weekday - 1) * 24))
	weekStart := today.Add(m)
	m, _ = time.ParseDuration(fmt.Sprintf("%dh", (7 - weekday + 1) * 24))
	weekEnd := today.Add(m)
	fmt.Println(weekStart, weekEnd)
	fmt.Println(weekStart.Format("2006-01-02") >= time.Now().Format("2006-01-02"))
	fmt.Println(weekEnd.Format("2006-01-02") >= time.Now().Format("2006-01-02"))

	//fmt.Println(int(weekday))
	//fmt.Println(int(time.Now().Weekday()))

	today_new, _ := time.Parse("2006-01-02 15:04:05", "2019-08-01 11:30:20")
	today_new2, _ := time.Parse("2006-01-02", "2019-08-02")
	fmt.Println(today_new2.Sub(today_new).Hours())
	fmt.Println(today_new.Format("2006-01-02"))
	add, _ := time.ParseDuration(fmt.Sprintf("%dh", 2))
	two_day_after := today_new.Add(add)
	fmt.Println(two_day_after.Format("2006-01-02"))

}
