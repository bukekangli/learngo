package container

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%d, len=%d cap=%d\n", s, len(s), cap(s))
}

func SlicepsTest() {
	fmt.Println("Creating Slice")
	var s []int // s == nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Delete elements from slice")

	s2  = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Poppping from front")

	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)

	fmt.Println("Poppping from back")

	back := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println(back)
}
