package main

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
	}
	for _, tt := range tests {
		if actual := Triangle(tt.a, tt.b); actual != tt.c {
			fmt.Printf("Triangle(%d, %d); "+
				"got:%d expect: %d; ", tt.a, tt.b, actual, tt.c)
		}
	}
}
