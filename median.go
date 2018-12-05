//median of two sorted arrays of equal lengths
package main

import (
	"fmt"
	"math"
)

func main(){
	a := []int{5,14,15,16,17,18} 
//a := []int{4,7,9,10}
	b := []int{10,11,12,13,14,15} 
//b := []int{3,5,6,8}

fmt.Println(median(a,b))
}

func median(a, b []int) float64{
	if len(a) ==1 {
		return float64(a[0] + b[0])/2
	}
	var median float64
	var v [2][5]int
	a_pad, b_pad := make([]int, len(a)+2), make([]int, len(b)+2)
        copy(a_pad[1:len(a_pad)-1], a)
	copy(b_pad[1:len(b_pad)-1] ,b)
	a_pad[len(a_pad)-1], b_pad[len(b_pad)-1] = math.MaxInt64, math.MaxInt64
	v[0][0], v[1][0] = 1, 1
	v[0][1], v[1][1] = len(a), len(b)
	v[0][2], v[1][2] = (len(a)+1)/2, (len(b)+1)/2
	v[0][3], v[1][3] = a_pad[v[0][2]], b_pad[v[1][2]]
	v[0][4], v[1][4] = a_pad[v[0][2]+1], b_pad[v[1][2]+1]
	taken := 0
	Loop:
		for v[0][0] <= v[0][1] && v[1][0] <= v[1][1] {
			v[0][2], v[1][2] = (v[0][0]+v[0][1])/2, (v[1][0]+v[1][1])/2
			v[0][3], v[1][3] = a_pad[v[0][2]], b_pad[v[1][2]]
                        v[0][4], v[1][4] = a_pad[v[0][2]+1], b_pad[v[1][2]+1]
			fmt.Println("taken", v[taken][3], "notTAken",v[1-taken][3])
			if v[taken][3] < v[1-taken][3]{
				taken = 1 - taken
			}else if v[taken][3] > v[1-taken][4]{
				v[taken][1] = v[taken][2]
				v[1-taken][0] = v[1-taken][2] + 1
			}else {
				break Loop
			}
		}
	count := (v[0][2] -1) + (v[1][2] -1) + 1
	fmt.Println(count)
	if count == (len(a)+len(b))/2{
		x := func(taken, pivot int)int{
			if taken == 0{
				return a_pad[pivot-1]
			}else {
				return b_pad[pivot-1]
			}
		}
		median = float64(v[taken][3] + compare("l", x(taken, v[taken][2]), v[1-taken][3]))/2
	}else {
		median = float64(v[taken][3] + compare("s", v[0][4], v[1][4]))/2
	}

	return median
}


func compare(i string, a int, b int) int{
	if (a > b && i == "l") || (a < b && i == "s"){
		return a
	}else {
		return b
	}
}
