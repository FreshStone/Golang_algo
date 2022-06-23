package main

import (
	"fmt"
	"heaputils" //package stored in $GOPATH/src/  but go runtime not searching it in $GOPATH/src
	//  solution --> go env -w GO111MODULE = off
)

func main() {
	var v1, v2 int
	pq := heaputils.PriorityQueue{}
	less := func(i, j int) bool {
		v1 = pq.Arr[i].(int) // type checking
		v2 = pq.Arr[j].(int) //this will panic if not checked against underlying concrete type
		return v1 < v2
	}
	pq.Push(3, less)
	pq.Push(4, less)
	pq.Push(2, less)
	top := pq.Pop(less)
	fmt.Println(top)
}
