package container

import (
	"fmt"
)

func lengthOfNonRepeatString(s string) int {
	// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func Test() {
	str := "abcadebcf"
	fmt.Println(lengthOfNonRepeatString(str))
	fmt.Println(lengthOfNonRepeatString("我是商力亢"))
	fmt.Println(lengthOfNonRepeatString("化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发"))

}
