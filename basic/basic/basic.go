package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"math"
)

func Triangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(100)
	s.Insert(1000)
	fmt.Println(s.Has(1))
}

func main() {
	fmt.Println(Triangle(3, 4))
	testSparse()
}
