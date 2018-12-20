package main

import (
	"fmt"
	"math"
)

var a = []int{7,10,4,3,20,15}
var heap_arr []int

func main(){
	for i, v := range a {
		insertintoheap(v, i)
	}
	Printheap()
	deletefromheap()
	Printheap()
	deletefromheap()
        Printheap()
	deletefromheap()
        Printheap()
}

func insertintoheap(v, l int){
	if l == len(heap_arr){
		heap_arr = append(heap_arr, v)
	}else {
		heap_arr[l] = v
	}
	if l == 0 {
		return
	}

	parent := int(math.Floor((float64(l)-1)/2))
	if v < heap_arr[parent]{
		heap_arr[l] = heap_arr[parent]
		insertintoheap(v, parent)
	}
	return
}

func Printheap(){
	level := int(math.Floor(math.Log2(float64(len(heap_arr))))) + 1
	l := level
	var first_space int //space before 1st character of a line
	var space int      //space between consecutive characters of a line
	line := 0
	for i:=0 ;i<len(heap_arr);i++{
		if i == int(math.Pow(2,float64(line)))-1{
			fmt.Println()//change line
			line += 1
			space = int(math.Pow(2, float64(l)))-1
			l -=1
			first_space = int(math.Pow(2, float64(l)))-1
			PrintSpaces(first_space)
			fmt.Printf("%v", heap_arr[i])
			continue
		}else{
			PrintSpaces(space)
			fmt.Printf("%v", heap_arr[i])
		}
	}
	fmt.Println()
}

func PrintSpaces(space int){
	for i:=0; i<space; i++{
		fmt.Printf(" ")
	}
}

func deletefromheap(){
	if len(heap_arr) == 0 {
		return
	}
	v := heap_arr[len(heap_arr)-1]
	heap_arr = heap_arr[:len(heap_arr)-1]
	var left_child, right_child int
	Loop :
	for i:=0; i<len(heap_arr);{
		left_child = 2*i + 1
		right_child = 2*i + 2
		if left_child > len(heap_arr) || len(heap_arr[left_child:]) == 0{
                	heap_arr[i] = v
			break Loop
	        }else if (len(heap_arr[right_child:]) == 0 || heap_arr[right_child] > heap_arr[left_child]){
                	heap_arr[i] = heap_arr[left_child]
               		i = left_child
        	}else {
                	heap_arr[i] = heap_arr[right_child]
                	i = right_child
        	}
	}
	return
}


