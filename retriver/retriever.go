package main

import (
	"fmt"
	"learngo/retriver/mock"
	"learngo/retriver/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

const url = "http://www.baidu.com"

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v.Contents)
	case *real.Retreiver:
		fmt.Println(v.UserAgent)
	}
	fmt.Println()
}

type Poster interface {
	Post(s string, form map[string]string) string
}

func post(p Poster) {
	p.Post("http://www.baiddu.com", map[string]string{
		"name":   "ccmouse",
		"course": "golan,",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(r RetrieverPoster) string {
	r.Post(url, map[string]string{
		"contents": "another fake baidu",
	})
	return r.Get(url)
}

func main() {
	var r Retriever
	retriver := &mock.Retriever{"this is a fake baidu.com"}
	inspect(retriver)
	r = retriver
	mockRetriever := r.(*mock.Retriever)
	fmt.Println(mockRetriever.Contents)
	r = &real.Retreiver{
		UserAgent: "Chrome/11.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//realRetriever := r(*real.Retreiver)

	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(retriver))

}
