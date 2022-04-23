package arithmetic

import (
	"fmt"
	"sync"
)

// 给定一个数组奇数排前、偶数排后
func SortOddEvenNum(l []int) []int {
	length := len(l)
	if length <= 1 {
		return l
	}
	oddIndex, evenIndex := length-1, 0
	findOdd, findEven := false, false
	for {
		// 从前往后找偶数
		for i := evenIndex; i < oddIndex && !findEven; i++ {
			if l[i]%2 == 0 {
				findEven = true
				evenIndex = i
			}
		}
		// 从后往前找奇数
		for j := oddIndex; j > evenIndex && !findOdd; j-- {
			if l[j]%2 == 1 {
				findOdd = true
				oddIndex = j
			}
		}
		if !findOdd || !findEven {
			break
		}
		l[evenIndex], l[oddIndex] = l[oddIndex], l[evenIndex]
		findOdd, findEven = false, false
		if oddIndex-evenIndex == 1 {
			break
		}
	}
	return l
}

func alterPrint() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-letter:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				number <- true
			}
		}
	}()
	wait.Add(1)
	go func() {
		c := 'A'
		for {
			select {
			case <-number:
				fmt.Println(c)
				c++
				fmt.Println(c)
				c++
				if c > 'Z' {
					wait.Done()
					return
				}
				letter <- true
			}
		}
	}()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2数之和
// 7->4->5 + 8->3->2 == 5->8->5
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2, 0)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 != nil {
		carry += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}
	dummy := &ListNode{
		Val:  carry % 10,
		Next: addTwoNumbers(l1, l2, carry/10),
	}
	return dummy
}
