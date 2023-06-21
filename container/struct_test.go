package container

import "testing"

type A struct {
	name string
}
type B struct {
	name string
}
type C struct {
	A
}
type D A
type E struct {
	D
}

func TestStructInherit(t *testing.T) {
	a := &A{}
	b := (*B)(a)
	c := C{*a}
	d := (*D)(a)
	e := E{*(*D)(a)}
	t.Logf("a: %v b: %v c: %v d: %v e: %v", a, b, c, d, e)
}
