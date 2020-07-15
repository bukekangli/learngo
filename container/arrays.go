package container

import (
	"fmt"
)

func printArray(arr *[4]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func ArrayTest() {
	var arrays [5]int
	arrays2 := [4]int{1, 2, 3, 4}
	arrays3 := [...]int{1, 3, 4, 4, 5, 5, 6}
	var grid [4][5]bool
	fmt.Println(arrays, arrays2, arrays3)
	fmt.Println(grid)

	for i := 0; i < len(arrays2); i++ {
		fmt.Println(arrays2[i])
	}

	printArray(&arrays2)
	fmt.Println(arrays2)
	a1 := []int{1, 2, 3}
	a2 := []int{1, 4}
	fmt.Println(a1, a2)
}
