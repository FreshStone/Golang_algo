package main

import "fmt"

func main(){
	//a := []int{3,7,6,8,4,2,1,9}
	//a := []int{8,1,4,2,6}
	a := []int{5,6,7,3,1}
	Quicksort(a,0,len(a)-1)
        fmt.Println(a)
}

func Quicksort(a []int, low, high int){
	var pivot_ind int
	if low <= high {
		pivot_ind = partition(a, low, high)
	}else{
		return
	}
	Quicksort(a, low, pivot_ind-1)
	Quicksort(a, pivot_ind+1, high)
}


func partition(a []int, low, high int) int{
	i := low
	pivot := a[high]
	for j := i; j <= high; j++{
		if a[j] <= pivot{
			/*
			s := a[i]
			a[i] = a[j]
			a[j] = s
			*/
			a[i], a[j] = a[j], a[i]
			i += 1
		}
	}
	return i-1
}

