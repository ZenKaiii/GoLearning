package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // go race 检测数据冲突
			//fmt.Printf("go %d\n",i)
			a[i]++
			//runtime.Gosched() go携程主动让出资源
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
