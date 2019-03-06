package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Print('h')
	}
}
func main() {
	readFile("abc.txt")
}
