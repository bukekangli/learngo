package main

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7}
	count := 5
	for i := 0; i < len(l); {
		end := i + count
		if end > len(l) {
			end = len(l)
		}
		t.Logf("%#v", l[i:end])
		i += count
	}
}

func TestLoop2(t *testing.T) {
loop:
	for {
		for {
			fmt.Println(1)
			break loop
		}
	}
}
