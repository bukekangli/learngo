package container

import (
	"fmt"
	"reflect"
	"testing"
)

func Remove(ptrArray interface{}, val interface{}) int {
	v := reflect.ValueOf(ptrArray)
	t := reflect.TypeOf(ptrArray)
	if t.Kind() != reflect.Ptr {
		panic("Must be a pointer of slice")
	}
	t = t.Elem()
	v = v.Elem()

	if t.Kind() != reflect.Slice {
		panic("Remove a non-slice type")
	}
	if !v.CanAddr() {
		panic("Array can not address")
	}

	vt := reflect.TypeOf(val)
	if t.Elem() != vt {
		panic("Elem and Val type not match")
	}

	var removed int
	for i, j := 0, 0; i < v.Len(); i++ {
		ei := v.Index(i)
		if reflect.DeepEqual(ei.Interface(), val) {
			removed += 1
		} else {
			v.Index(j).Set(ei)
			j += 1
		}
	}
	v.SetLen(v.Len() - removed)

	return removed
}

func editSlice(nums []int) {
	Remove(&nums, 5)
	fmt.Println(nums)
}

func TestSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	editSlice(nums)
	t.Logf("nums: %#v", nums)
}
