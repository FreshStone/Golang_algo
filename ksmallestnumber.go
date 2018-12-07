package main

import (
	"fmt"
)

func main(){
	a := []int{5,3,2,6,1,10}
	k := 2
	fmt.Println(findsmallest(a, k))
}

func findsmallest(a []int, k int)int{
	if k == 0 || k > len(a){
		return 0
	}
	b := make([]int, k)
	for i, v := range a{
		if i < k{
			insertintoheap(b, v, i, i)
		}else if a[i] < b[0]{
			insertintoheap(b, v, 0, k)
		}else {
			continue
		}
	}
	return b[0]
}

func insertintoheap(b []int, v, i, l int){
	b[i] = v
	if i == l {
		parent := (i-1)/2
		if v > b[parent]{
			b[i] = b[parent]
			insertintoheap(b, v, parent, parent)
		}
	}else{
		leftchild := 2*i + 1 ; rightchild := 2*i + 2
		c := max(b, leftchild, rightchild)
		if c != -1 && v < b[c]{
			b[i] = b[c]
			insertintoheap(b, v, c, l)
		}
	}
	return
}

func max(b []int, l, r int) int{
	if l > len(b)-1{
		return -1
	}else if r > len(b)-1 || b[l] > b[r]{
		return l
	}else {
		return r
	}
}
