package test

import (
	"fmt"
	"os"
	"testing"
)

func TestFileAppend(t *testing.T) {
	for i := 0; i < 10; i++ {
		f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		f.WriteString(fmt.Sprintf("%d\n", i))
	}
}
