package test

import (
	"reflect"
	"testing"
)

func Test_ReflectGetVariableName(t *testing.T) {
	var slk string
	typeOf := reflect.TypeOf(slk)
	t.Logf("%#v", typeOf)
	valueOf := reflect.ValueOf(slk)
	t.Logf("%#v", valueOf)

}
