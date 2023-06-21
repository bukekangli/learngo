package test

import (
	"sync"
	"testing"
	"time"
)

func TestForBreak(t *testing.T) {
	ch := make(chan bool, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			val, ok := <-ch
			if !ok {
				t.Logf("break")
				break
			}
			t.Logf("val: %t", val)
			time.Sleep(1 * time.Second)
		}

	}()
	for i := 0; i < 3; i++ {
		ch <- true
	}
	close(ch)
	wg.Wait()
}
