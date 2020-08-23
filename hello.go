package main

import (
	"fmt"
	"learngo/arithmetic"
)

func main() {
	fmt.Printf("hello, world\n")
	l := []int{4343, 1, 7, 13, 4, 234, 3234, 5, 6, 3, 8, 9, 11, 321, 224, 999, 123, 444, 333}
	//fmt.Println(arithmetic.SortOddEvenNum([]int{1,2,3,4,5,6,7,8,9,11,321,224,999,123,444,333}))
	//arithmetic.InsertSort(l)
	//arithmetic.BubbleSort(l)
	//arithmetic.SelectSort(l)
	//fmt.Println(arithmetic.MyMergeSort(l))
	arithmetic.MergeSort(l)
	fmt.Println(l)
	//fmt.Println(l[:3])
}
