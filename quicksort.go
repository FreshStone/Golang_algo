package main

import "fmt"

func main(){
	a := []int{3,7,6,8,4,2,1,9}
	quicksort(a,0,len(a)-1)
        fmt.Println(a)
}

func quicksort(a []int, low, high int){
	var pivot_ind int
	if low <= high {
		pivot_ind = partition(a, low, high)
	}else{
		return
	}
	quicksort(a, low, pivot_ind-1)
	quicksort(a, pivot_ind+1, high)
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

