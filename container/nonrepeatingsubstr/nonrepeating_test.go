package nonrepeatingsubstr

import (
	"fmt"
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s  string
		an int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发", 7},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatString(tt.s); actual != tt.an {
			fmt.Printf("lengthOfNonRepeatString(%s) got %d;"+
				"expect %d", tt.s, actual, tt.an)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发"
	an := 7
	for i := 0; i < b.N; i++ {

		if actual := lengthOfNonRepeatString(s); actual != an {
			fmt.Printf("lengthOfNonRepeatString(%s) got %d;"+
				"expect %d", s, actual, an)
		}
	}
}
