package main

import (
	"fmt"
	"math/rand"
	"time"
)

//标记字符 与 只能写入的通道
func producer(header string, channel chan<- string)  {
	for  {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string)  {
	for {
		//message := <-channel
		fmt.Println(<-channel)
	}
}


//func main() {
//	channel := make(chan string)
//	go producer("dog",channel)
//	go producer("cat",channel)
//	customer(channel)
//}
