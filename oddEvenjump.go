package main

import "fmt"

type node struct{
	id int
	is_odd_good, is_even_good bool
	left, right *node
}

var a = []int{3,4,2,2,5,6,4,8}

func main(){
	l := len(a)
	good_jumps := 1
	jump_store := make([]*node, l)
	root := &node{id:l-1,is_odd_good:true, is_even_good:true}
	jump_store[l-1] = root
	for i := l-2; i>=0; i--{
        	jump_store[i] = &node{id:i}
		odd, even := insert(root, i)
        	if odd > i && jump_store[odd].is_even_good{
                	jump_store[i].is_odd_good = true
                	good_jumps++
        	}

		if even > i && jump_store[even].is_odd_good{
                	jump_store[i].is_even_good = true
		}
	}
	fmt.Println(good_jumps)
}

func insert(root *node, i int)(int, int){
	var odd, even int
	equal := -1
	for{
		if a[i] > a[root.id]{
			even = root.id
			if root.right == nil{
				root.right = &node{id:i}
				break
			}else{
				root = root.right
			}
		}else{
			if a[i] < a[root.id]{
				odd = root.id
			}else{
				equal = root.id
			}

			if root.left == nil{
                                root.left = &node{id:i}
                                break
                        }else{
                                root = root.left
                        }
		}
	}

	if equal == -1{
		return odd, even
	}else{
		return equal, equal
	}
}
