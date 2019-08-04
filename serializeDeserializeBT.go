package main

import (
	"fmt"
	"strconv"
)

type node struct{
	val int
	left, right *node
}

type codec struct{}

func (c codec) serialize(n *node)string{
	if n == nil{
		return ""
	}
	return "(" + c.serialize(n.left) + strconv.Itoa(n.val) + c.serialize(n.right) + ")"
}

func (c codec) deserialize(s string)(*node){ //serialized in inorder format
	if len(s) == 0{
		return nil
	}
	n, _ := decode(s, 1)
	return n
}

func decode(s string, i int)(*node, int){
	root := &node{}
	if s[i] == 40{ //open bracket
		root.left, i = decode(s, i+1)
	}

	tmp := i
	for ;s[tmp] != 40 && s[tmp] != 41;{
		tmp += 1
	}
	root.val, _ = strconv.Atoi(s[i:tmp])
	if s[tmp] == 41{ //close bracket
		return root, tmp+1
	}

	root.right, i = decode(s, tmp+1)
	return root, i+1
}



func main(){
	var c codec
	n := &node{val: 1}
	n.left = &node{val: 2}
	n.left.left = &node{val: 3}
	n.left.left.left = &node{val: 4}
	n.left.left.left.left =  &node{val: 5}
	n.left.right =  &node{val: 6}
	n.left.right.right =  &node{val: 7}
	n.left.right.right.right =  &node{val: 8}
	n.right =  &node{val: 9}
	n.right.left =  &node{val: 10}
	n.right.left.left =  &node{val: 11}
	n.right.left.right =  &node{val: 12}
	n.right.left.right.left =  &node{val: 13}
	n.right.left.right.right =  &node{val: 14}
	n.right.right =  &node{val: 15}

	s := c.serialize(n)
	fmt.Println(s) //(((((5)4)3)2(6(7(8))))1(((11)10((13)12(14)))9(15)))
	n = c.deserialize(s)
	printinorder(n)
	fmt.Println()
}

func printinorder(n *node){
	if n == nil{
		return
	}
	printinorder(n.left)
	fmt.Print(" ", n.val)
	printinorder(n.right)
}
