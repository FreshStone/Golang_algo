package main

import (
	"fmt"
	"math"
)

var f = 5
var e = 3
var dp = make([][]uint16, e+1)

func main(){
	for i, _ := range dp{
		dp[i] = make([]uint16, f+1)
	}

	for i := 1; i < e+1; i++{
		for j :=1; j<f+1; j++{
			if i == 1 {
				dp[i][j] = uint16(j)
			}else{
				dp[i][j] = min(i, j)
			}
		}
	}
	fmt.Println(dp[e][f])
}

func min(i, j int)uint16{
	m := uint16(math.Pow(2,16) - 1)
	for k:=1; k<=j; k++{
		n := 1 + max(dp[i-1][k-1], dp[i][j-k])
//		fmt.Println(n)
		if m > n {
			m = n
		}
	}
//	fmt.Println("m", m)
	return m
}

func max(a, b uint16) uint16{
	if a > b {
		return a
	}else{
		return b
	}
}





