package main

import (
	"fmt"
	"learngo/arithmetic"
)

func main() {
	fmt.Printf("hello, world\n")
	l := []int{4343, 1, 2, 3, 4, 234, 3234, 5, 6, 7, 8, 9, 11, 321, 224, 999, 123, 444, 333}
	//fmt.Println(arithmetic.SortOddEvenNum([]int{1,2,3,4,5,6,7,8,9,11,321,224,999,123,444,333}))
	//arithmetic.InsertSort(l)
	//arithmetic.BubbleSort(l)
	arithmetic.SelectSort(l)
	fmt.Println(l)
}
