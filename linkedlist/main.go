package main

import "fmt"

type node struct{
	val int
	next *node
}

var r node

func main(){
	for _, v := range []int{1,2,3,4,5}{
		insert(&r, v)
	}
	find(&r)
}

func insert(root *node, a int){
        var n node
        if (*root).next == nil {
                (*root).next = &n
                n.val = a
                return
        }
        insert((*root).next, a)
}

func find(r *node){
        if (*r).next == nil {
                return
        }
        n := (*r).next
        fmt.Println(n.val)
        find((*r).next)
}

