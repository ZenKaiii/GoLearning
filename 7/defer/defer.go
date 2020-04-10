package main

import (
	"GoLearning/6/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	for i:=0; i<100 ;i++  {
		defer fmt.Println(i)
		if i == 30{
			panic("too many")
		}
	}
}

func Num(i int)int  {
	//f
	return i
}

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}

	//if file, err := os.Create(filename); err!=nil{
	//	panic(err)
	//}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i:=0;i<20 ;i++  {
		fmt.Fprintln(writer,f())
	}
}

func main() {
	writeFile("hello.txt")
	//tryDefer()
}
