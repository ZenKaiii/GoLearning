package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func plus(a int, b int) int {
	return a + b
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
	name string
	age  int
}

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't do")
	}
	return arg + 3, nil
}

func f(form string) {
	for i := 0; i < 3; i++ {
		fmt.Println(form, ":", i)
	}
}

func worker(done chan bool)  {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string)  {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string)  {
	msg := <-pings
	pongs <- msg
}

func main() {

	//非阻塞通道


	//超时处理
	//c1 := make(chan string)
	//go func() {
	//	time.Sleep(time.Second * 1)
	//	c1 <- "result 1"
	//}()
	//
	//select {
	//case res:= <-c1:
	//	fmt.Println(res)
	//case <- time.After(time.Second * 2):
	//	fmt.Println("time out")
	//
	//}


	// 通道选择器
	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(time.Second *1)
	//	c1 <- "one"
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second * 1)
	//	c2 <- "two"
	//}()
	//
	//for i:=0;i<2 ;i++  {
	//	select {
	//	case msg1:= <-c1:
	//		fmt.Println(msg1)
	//	case msg2:= <-c2:
	//		fmt.Println(msg2)
	//
	//	}
	//}

	//done := make(chan bool, 1)
	//go worker(done)
	//<- done
	//messages := make(chan string, 2)
	//
	//messages <- "hello"
	//messages <- "world"
	//messages <- "world"
	//
	//
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

	//messages := make(chan string)
	//
	//go func() {
	//	messages <- "ping"
	//}()
	//
	//msg := <- messages
	//fmt.Println(msg)

	//f("direct")
	//
	//go f("goroutine")
	//
	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")
	//
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("done")

	//for _,i := range []int{7,42} {
	//	if r,e := f1(i); e != nil {
	//		fmt.Println("failed",e)
	//	}else{
	//		fmt.Println("success",r)
	//	}
	//}

	//r := rect{width: 3, height: 4}
	//c := circle{radius: 5}
	//
	//measure(r)
	//measure(c)

	//f := rect{5,10}
	//fmt.Println(f.area())

	//f := "hello"
	//fmt.Println(f)
	//
	//for i:=1; i<=3; i++{
	//	fmt.Println(i)
	//}

	//var a [5]int
	//a[4] = 100
	//fmt.Println(a[0])

	//b := [5]int{1,2,3,4,5}
	//for i:=0;i<len(b) ;i++  {
	//	fmt.Println(b[i])
	//}
	//fmt.Println(b)

	//var c [2][3]int
	//for i:=0;i<2 ;i++  {
	//	for j:=0;j<3 ;j++  {
	//		c[i][j] = i+j
	//	}
	//}
	//fmt.Println(len(c))
	//fmt.Println(len(c[0]))
	//fmt.Println(c)

	//s := make([]string,3)
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	//fmt.Println(s)
	//s = append(s, "d")
	//fmt.Println(s)
	//
	//c := make([]string,len(s))
	//copy(c,s)
	//fmt.Println(c)

	//twoD := make([][]int, 3)
	//fmt.Println(twoD)
	//for i:=0;i<len(twoD) ;i++  {
	//	innerLength := i+1
	//	twoD[i] = make([]int,innerLength)
	//	for j:=0;j<innerLength ;j++  {
	//		twoD[i][j] = i+j
	//	}
	//}
	//fmt.Println(twoD)

	//m := make(map[string]int)
	//m["k1"] = 7
	//m["james"] = 23
	//fmt.Println(m)
	//fmt.Println(m["james"])
	//m["ww"] = 3
	//fmt.Println(m)
	//delete(m, "ww")
	//fmt.Println(m)
	//
	//_, prs := m["k3"]
	//fmt.Println(prs)
	//
	//n := map[string]int{"a":1}
	//fmt.Println(n)

	//nums := []int{1,2,3}
	//sum := 0
	//for _,num := range (nums){
	//	sum += num
	//}
	//fmt.Println(sum)
	//
	//for i,c:=range ("Aa"){
	//	fmt.Println(i,c)
	//}
	//
	//kvs := map[string]int{"abc":1,"james":23}
	//for k,v := range kvs{
	//	fmt.Println(k,v)
	// }
	//fmt.Println(plus(1,2))
	//nums := []int{1,2,3,4}
	//sum(nums...)

	//nextInt := intSeq()
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//
	//nextInts := intSeq()
	//fmt.Println(nextInts())

	//fmt.Println(fact(3))

	//i := 1
	//zeroptr(&i)
	//fmt.Println(i)
	//fmt.Println(&i)

	//fmt.Println(person{"ann",12})
	//s := person{
	//	name: "asd",
	//	age:  12,
	//}
	//fmt.Println(s)
	//s.age = 123
	//fmt.Println(s)
}
