package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		res := b
		a, b = b, a+b
		fmt.Println(res)
		return res
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%s\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}

func main() {
	f := Fibonacci()
	PrintFileContents(f)
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
}
