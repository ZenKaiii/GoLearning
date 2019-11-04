package main

import (
	"GoLearning/retriever/mock"
	"GoLearning/retriever/reall"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is imooc"} //拷贝
	fmt.Printf("%T %v\n",r,r)
	r = &reall.Retriever{ //传指针
		UserAgner:"Mozilla/5.0",
		Timeout:time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)

	//fmt.Println(download(r))
}
