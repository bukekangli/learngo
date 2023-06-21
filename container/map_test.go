package container

import (
	"fmt"
	"sync"
	"testing"
)

func MapTest(t *testing.T) {
	m := map[string]string{
		"name": "slk",
		"age":  "27",
	}
	m["school"] = "sy"

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int //  m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	if name2, ok := m["name2"]; ok == true {
		fmt.Println(name2)
	} else {
		fmt.Printf("key %s does not exists", name2)
	}

	fmt.Println("Delete values")
	delete(m, "name")
}

func Test_ConcurrencyRead(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for {
				for j := 0; j < 10; j++ {
					_ = m[j]
				}
			}
		}()
	}
	wg.Wait()
}
