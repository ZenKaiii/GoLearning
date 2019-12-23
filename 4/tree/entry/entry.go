package main

import (
	"GoLearning/4/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	fmt.Println(root)
	root.Print()
}
