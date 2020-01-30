package main

import (
	"fmt"
)

func grade(score int) string{
	g := ""
	switch  {
	case score < 0 || score >100:
		panic(fmt.Sprintf("wrong : %d",score))
	case score <60:
		g = "F"
	case score<80:
		g = "C"
	case score <90:
		g = "B"
	case score <=100:
		g = "A"
	default:
		panic("wrong")

	}
	return g
}

//func main() {
//	const filename = "2/abc.txt"
//	//contents, err := ioutil.ReadFile(filename)
//	//if err != nil{
//	//	fmt.Println(err)
//	//} else {
//	//	fmt.Printf("%s\n", contents)
//	//}
//	if contents, err := ioutil.ReadFile(filename);err!=nil{
//		fmt.Println(err)
//	}else {
//		fmt.Printf("%s\n", contents)
//	}
//
//}
