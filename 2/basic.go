package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
var ( //全局变量
	aa = 33
	bb = true
	ss = "kk"
)


func variableZeroValue(){
	var a int
	var b string
	fmt.Printf("%d %q \n",a,b)
}

func variableInitialValue()  {
	var a, b int = 3 , 4
	var s string = "aaa"
	fmt.Println(a,b,s)
}

func variableTypeDeduction() {
	var a, b = 3, 4
	var s = "sss"
	fmt.Println(a, b, s)
}

func variableShorter(){
	a,b,c,d := 3,4,true,"def"
	fmt.Println(a,b,c,d)
}

func euler()  {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	//fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f \n",cmplx.Exp(1i*math.Pi)+1)

}

func triangle(){
	a,b := 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a,b = 3,4
	c := int(math.Sqrt((a*a + b*b)))
	fmt.Println(c)
}

func enums()  {
	const (
		cpp = iota //自增值
		java
		python
	)
	fmt.Println(cpp,java,python)
}

//func main() {
//	fmt.Println("hello world")
//	variableZeroValue()
//	variableInitialValue()
//	variableTypeDeduction()
//	variableShorter()
//	fmt.Println(aa,ss,bb)
//	euler()
//	triangle()
//	consts()
//	enums()
//}
