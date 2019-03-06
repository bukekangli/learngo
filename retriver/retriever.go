package main

import (
	"fmt"
	"learngo/retriver/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake baidu.com"}
	fmt.Println(download(r))
}

