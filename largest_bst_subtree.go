package main

import "fmt"

type node struct{
	val int
	left, right *node
}

func main(){
	root := &node{val: 4}
	root.left = &node{val:1}
	root.right = &node{val:8}
	root.right.left = &node{val: 6}
	root.right.right = &node{val: 9}
	root.right.left.left = &node{val: 3}
	root.right.left.right = &node{val: 7}
	root.right.right.right = &node{val: 10}
	root.right.right.right.right = &node{val: 11}
	root.right.right.right.right.right = &node{val: 12}


	fmt.Println(lbst(root)) // root != nil
}

func lbst(root *node)(int, bool, int, int){
	if root.left == nil && root.right == nil{
		return 1, true, root.val, root.val
	}

	var l, r, ltl, ltr, rtl, rtr int
	var ok1, ok2 bool
	if root.left != nil{
		l, ok1, ltl, ltr = lbst(root.left)
	}else{
		l, ok1, ltl, ltr = 0, true, root.val, root.val
	}


	if root.right != nil{
		r, ok2, rtl, rtr = lbst(root.right)
	}else{
		r, ok2, rtl, rtr = 0, true, root.val, root.val
	}

	if ok1 && ok2 && ltr <= root.val && root.val <= rtl{
			return l+r+1, true, ltl, rtr
	}

	return max(l, r), false, -1, -1
}

func max(a...int)int{
	max := a[0]
	for i := 1; i < len(a); i++{
		if a[i] > max{
			max = a[i]
		}
	}
	return max
}

func maxpath(root *node)(int, int){
	if root == nil{
		return -1, -1
	}

	l, li := maxpath(root.left)
	r, ri := maxpath(root.right)

	return 1+max(l, r), max(li, ri, l+r+2)
}
