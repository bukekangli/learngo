package main

import (
	"fmt"
	"time"
)

func ReturnFunc() func() {
	// 测试闭包 defer执行顺序
	start_ts := time.Now()
	return func() {
		//fmt.Println(a*a)
		fmt.Printf("run cost: %s", time.Now().Sub(start_ts))
	}
}
func simple() {
	fmt.Println("haha ")
}
func main() {
	defer ReturnFunc()()
	time.Sleep(2 * time.Second)
	//f := fib.Fibonacci()
	//fib.PrintFileContents(f)
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
}
