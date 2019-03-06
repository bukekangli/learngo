package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)

	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int)  (int, int){
	return b, a
}

func main() {
	fmt.Println(div(10, 3))
	fmt.Println(eval(11, 3, "/"))
	if result, err := eval(11, 3, "x"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(apply(pow, 11, 2))
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 11, 2))
	fmt.Println(sum(1, 2, 3, 4))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a,b)
}
