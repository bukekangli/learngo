package container

import (
	"fmt"
	"unicode/utf8"
)

func StringTest() {
	s := "商力亢最帅" // utf-8
	fmt.Println("len = ", len(s))
	//for _, b := range []byte(s) {
	//	fmt.Printf("%X ", b)
	//}
	for i, b := range s { // b is a rune
		fmt.Printf("(%d %X)", i, b)
	}
	fmt.Println()

	fmt.Println("Rune count ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
