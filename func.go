package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a/b
	default:
		panic("buzhidao "+op)

	}
}

func div(a,b int) (q,r int)  {
	return a/b , a%b
	//q = a/b
	//r = a%b
	//return
}

func apply(op func(int,int) int, a,b int) int {
	p:=reflect.ValueOf(op).Pointer()
	opname := runtime.FuncForPC(p).Name()
	fmt.Printf("%s \n", opname)
	return op(a,b)
}

func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func sum(numbers ...int) int{
	s := 0
	for i:= range numbers{
		s += numbers[i]
	}
	return s
}

func swap(a,b *int)  {
	*b,*a = *a,*b
}

//func main() {
//	fmt.Println(eval(3,4,"*"))
//	q,r := div(13,3)
//	fmt.Println(q,r)
//	fmt.Println(apply(pow,3,4))
//	fmt.Println(sum(1,2,3,4,5))
//
//	a,b := 3,4
//	swap(&a, &b)
//	fmt.Println(a,b)
//}
