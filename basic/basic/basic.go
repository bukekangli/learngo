package main

import (
	"fmt"
	"math"
)

func Triangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	fmt.Println(Triangle(3, 4))
}
