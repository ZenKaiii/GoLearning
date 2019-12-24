package main

import (
	real2 "GoLearning/5/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string)  string
}


func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	r := &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)
	//fmt.Println(download(r))
}
