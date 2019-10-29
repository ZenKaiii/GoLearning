package main

import "fmt"

func printArray(arr [5]int)  {
	arr[0] = 100 //值类型，所以不变原来的
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

//func main() {
//	var arr1 [5]int
//	arr2 := [3]int{1,3,5}
//	arr3 := [...]int{2,3,4,5,6}
//	var grid [4][5]int //四个长度为5的int数组
//	fmt.Println(arr1,arr2,arr3)
//	fmt.Println(grid)
//
//	//for i:=0; i<len(arr3); i++{
//	//	fmt.Println(arr3[i])
//	//}
//	printArray(arr3)
//}
