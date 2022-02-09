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
			case <- letter:
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
			case <- number:
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