package main

import "fmt"

type node struct{
	val int
	left, right *node
}

func main(){
	root := &node{val:3} //no duplicates in tree
	root.left = &node{val:6}
	root.right = &node{val:8}
	root.left.left = &node{val:2}
	root.left.right = &node{val:11}
	root.right.right = &node{val:13}
	root.right.right.left = &node{val:7}
	root.left.right.left = &node{val:9}
	root.left.right.right = &node{val:5}
//	printtree(root)
// 	fmt.Println()
	fmt.Println(lca(root, 9, 5))
}

func printtree(root *node){
	if root == nil{
		return
	}
	printtree(root.left)
	fmt.Printf("%d ", root.val)
	printtree(root.right)
	return
}

func lca(root *node, m, n int) int{
	var rs int
	if root == nil{
		return -2
	}
	if root.val == m || root.val == n{
		if lca(root.left, m, n) == -2 && lca(root.right, m, n) == -2{
			return -1
		}else{
			return root.val
		}
	}

	rs = lca(root.left, m, n)
        if rs == -2{
                return lca(root.right, m, n)
        }else if rs == -1{
		if lca(root.right, m, n) == -1{
			return root.val
		}else {
			return -1
		}
        }else{
		return rs
	}
}
