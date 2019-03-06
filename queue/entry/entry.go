package main

import (
	"fmt"
	"learngo/queue"
)

func main()  {
	q := queue.Queue{1}
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}