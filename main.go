package main

import (
	"fmt"
	"learngo/arithmetic"
)

func main() {
	//container.StringTest()
	//container.MapTest()
	//container.SlicepsTest()
	l1 := &arithmetic.ListNode{Val: 2}
	l1.Next = &arithmetic.ListNode{Val: 4}
	l1.Next.Next = &arithmetic.ListNode{Val: 3}
	l2 := &arithmetic.ListNode{Val: 5}
	l2.Next = &arithmetic.ListNode{Val: 6}
	l2.Next.Next = &arithmetic.ListNode{Val: 4}
	res := arithmetic.AddTwoNumbers(l1, l2)
	fmt.Println(res)
}
