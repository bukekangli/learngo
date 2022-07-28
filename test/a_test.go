package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

func TestTime(t *testing.T) {
	for _, d := range []string{
		"2022-07-18",
		"2022-07-19",
		"2022-07-20",
		"2022-07-21",
		"2022-07-22",
		"2022-07-23",
		"2022-07-24",
	} {

		t1, _ := time.Parse("2006-01-02", d)
		t.Logf(t1.Format("Monday"))
	}
}

func f1(s string) int {
	return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
	return strings.Count(s, "") - 1
}

func f3(s string) int {
	return len([]rune(s))
}

func f4(s string) int {
	return utf8.RuneCountInString(s)
}

func TestLen(t *testing.T) {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println([]byte(s))
	fmt.Println(f1(s))
	fmt.Println(f2(s))
	fmt.Println(f3(s))
	fmt.Println(f4(s))
}

func TestLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := 0
	loop:
		for {
			for {
				j++
				if j == 1 {
					break loop
				}
				t.Logf("can not print me")
			}
			t.Logf("can not print me")
		}
		t.Logf("%d", i)
	}
}
