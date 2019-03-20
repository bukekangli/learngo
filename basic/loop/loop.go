package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
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
	readFile("basic/loop/abc.txt")
	s := `abc
	def
	ghk
	lm""`
	printFileContents(strings.NewReader(s))

}
