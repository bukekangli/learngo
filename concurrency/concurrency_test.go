package concurrency

import (
	"testing"
	"time"
)

func TestConcurrentReadMap(t *testing.T) {
	m := make(map[int]bool)
	for i := 0; i < 10; i++ {
		m[i] = true
	}
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for {
	//			_ = m[1]
	//		}
	//	}()
	//}
	go func() {
		for i := 0; i < 2; i++ {
			go func() {
				for {
					m[11] = true
				}
			}()
		}
	}()
	time.Sleep(1 * time.Minute)
}
