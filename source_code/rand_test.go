package source_code

import (
	"math/rand"
	"testing"
)

func BenchmarkRand63n(t *testing.B) {
	for i := 0; i < t.N; i++ {
		//t.Logf("randInt63n: %d", rand.Int63n(10000))
	}
	for {
		if rand.Int63n(10000) == 1 {
			t.Logf("run break")
			break
		}
	}
}
