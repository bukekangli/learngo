package container

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func SliceTest() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("array[2:6] = ", array[2:6])
	fmt.Println("array[:6] = ", array[:6])
	s1 := array[2:]
	fmt.Println(len(s1))
	fmt.Println("s1 = ", s1)
	updateSlice(s1)
	fmt.Printf("After updateSlice(s1)")
	fmt.Println("s1 = ", s1)
	fmt.Println("array = ", array)
	s2 := array[:]

	fmt.Println("Reslice")
	fmt.Println("s2 = ", s2)
	s2 = s2[:5]
	fmt.Println("s2 = ", s2)
	s2 = s2[2:]
	fmt.Println("s2 = ", s2)

	fmt.Println("Extending slice")

	array[0], array[2] = 0, 2
	s1 = array[2:6]
	s2 = s1[3:5]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	fmt.Printf("s1 = %v lend(s1) = %d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v lend(s2) = %d cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	fmt.Println("array = ", array)

	s1 = array[:len(array)-1]
	fmt.Printf("s1 = %v lend(s1) = %d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
