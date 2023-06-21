package test

import (
	"testing"
)

var global interface{}

// interface引起的内存泄漏
// 1.go test -run none -bench Interface -benchmem -memprofile mem.out interface_test.go
// 2.go tool pprof -alloc_space -flat mem.out
//
//	2.1 top
//	2.2 list BenchmarkInterface
func BenchmarkInterface(b *testing.B) {
	var local interface{}
	for i := 0; i < b.N; i++ {
		// assign value to interface{}
		local = calculate(i)
	}
	global = local
}

func BenchmarkInterfaceV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// assign value to interface{}
		_ = calculate(i)
	}
}

var globalValue values

func BenchmarkInterfaceV3(b *testing.B) {
	var localValue values
	for i := 0; i < b.N; i++ {
		// assign value to interface{}
		localValue = calculate(i)
	}
	globalValue = localValue
}

// values is bigger than single machine word.
type values struct {
	value  int
	double int
	triple int
}

func calculate(i int) values {
	return values{
		value:  i,
		double: i * 2,
		triple: i * 3,
	}
}
