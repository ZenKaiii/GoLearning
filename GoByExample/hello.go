package main

import "fmt"

func plus(a int, b int) int  {
	return a+b
}

func sum(nums...int) {
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

func fact(n int) int {
	if n==0 {
		return 1
	}
	return n * fact(n-1)
}

func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
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

	fmt.Println(fact(3))
}
