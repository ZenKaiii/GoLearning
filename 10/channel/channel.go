package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo()  {
	var channels [10]chan<- int //send-only
	for i := 0;  i < 10; i++ {
		channels[i] = createWorker(i)
		//go worker(i, channels[i])
	}

	for i:=0; i <10; i++ {
		channels[i] <- 'a'+i
	}
	//c := make(chan int)
	//go func() {
	//	for {
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//}()
	//c <- 1
	//c <- 2
	time.Sleep(time.Millisecond)
}

func bufferedChannel()  {
	c := make(chan int, 3) //缓冲区
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c := make(chan int, 3) //缓冲区
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
