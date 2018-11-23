package main

import "fmt"

func main() {
	arr := []int{1, 14, 15, 2, 6, 7, 2, 9, 3}
	sorted_arr := mergesort(arr)
	fmt.Println(sorted_arr)
}

func mergesort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	l := mergesort(a[:(len(a)+1)/2])
	r := mergesort(a[(len(a)+1)/2:])
	m := merge(l, r)
	return m
}

func merge(a, b []int) []int {
	var m  []int
	i, j := 0, 0
	for (len(a[i:]) > 0 && len(b[j:]) > 0) {
	if a[i] < b[j] {
		m = append(m,a[i])
		i = i+1
	} else {
		m = append(m,b[j])
		j = j+1
	}
}

for (len(a[i:]) > 0) {
	m = append(m, a[i])
	i = i+1
}

for (len(b[j:]) > 0){
	m = append(m, b[j])
	j = j+1
}

return m
}


