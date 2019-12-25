package main

import (
	"GoLearning/5/retriever/mock"
	real2 "GoLearning/5/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string)  string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}


func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func post(p Poster)  {
	p.Post("https://www.baidu.com", map[string]string{
		"name":"zk",
		"course":"golang",
	})
}

func session(s RetrieverPoster) string {
	s.Post("https://www.baidu.com", map[string]string{
		"contents":"another fake imooc",
	})
	return s.Get("https://www.baidu.com")
}

func main() {
	r := real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)

	s := &mock.Retriever{"this is baidu"}

	fmt.Println(s)

	fmt.Println(session(s))
	//fmt.Println(download(r))
}
