package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type MyType int64

func TestReflect1(t *testing.T) {
	var a float64 = 1.0
	t.Logf("type: %s", reflect.TypeOf(a))
	var b MyType = 10
	t.Logf("type: %s", reflect.TypeOf(b))
	val := reflect.ValueOf(b)
	t.Logf("val.kind: %s", val.Kind())
	t.Logf("val.Type: %s", val.Type())
	val2 := val.Interface().(MyType)
	t.Logf("val.Interface().(MyType): %d", val2)
}

func TestReflect2(t *testing.T) {
	type T struct {
		A int
		B string
	}
	tObj := T{23, "skidoo"}
	s := reflect.ValueOf(&tObj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
