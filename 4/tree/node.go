package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode{
	return &treeNode{value : value}
}

func main() {
	var root treeNode
	fmt.Println(root)
}
