package test

import (
	"context"
	"testing"
)

func TestCtx1(t *testing.T) {
	ctx := context.Background()
	ctxWithCancel1, cancel1 := context.WithCancel(ctx)
	select {
	case <-ctxWithCancel1.Done():
	default:
		t.Logf("pass ---")

	}
	ctxWithCancel2, cancel2 := context.WithCancel(ctxWithCancel1)
	cancel1()
	cancel2()
	t.Logf("%#v %#v %#v %#v", ctxWithCancel1, cancel1, ctxWithCancel2, cancel2)
}
