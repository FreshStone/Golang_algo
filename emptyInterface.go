package main

import "fmt"

func main() {
	ans := sum([]int{1, 2}, []int{3, 4})
	fmt.Printf("type- %T, val: %v\n", ans, ans)
}

func sum(a, b interface{}) interface{} { //empty interface as i/o
	//return a + b //invalid operation: operator + not defined on a empty interface
	i, ok1 := a.([]int)
	j, ok2 := b.([]int)
	if ok1 && ok2 {
		return i[0] + j[0]
	}
	return nil
}
