package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func (node *treeNode) traverse() {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) { //传引用
	node.value = value
}

func main() {
	var root  = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)
	//root.right.left.print()
	//fmt.Println(root)
	//nodes := []treeNode{
	//	{value:3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	//root.print()
	root.traverse()
}
