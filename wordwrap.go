package main

import (
	"fmt"
	"math"
)
var a = []int{6,1,3,1}; var w = 8 //len of consecutive words
//var a = []int{3,2,2,5}
var index = make([]int, len(a))
//var w = 5 //max width of each line


func main(){
	dp := make([]int, len(a)+1)
	var n, min_cost int
	dp[0] = 0;
	for i := 1; i <= len(a);i++{
		n = a[i-1]
		min_cost = dp[i-1] + int(math.Pow(float64(w-n), 3))
		j := i-1
		index[i-1] = j
		Loop :
			for ; j >=1; j-- {
				n = n + a[j-1] +1

				if n > w  {
					break Loop
				}
				if min_cost >  dp[j-1] + int(math.Pow(float64(w-n), 3)){
					min_cost = dp[j-1] + int(math.Pow(float64(w-n), 3))
				//	index[i-1] = j
					index[i-1] = j-1
				}
			}
		dp[i] = min_cost
	}
	fmt.Println("min_cost-", dp[len(a)])
	printlines()
}

func printlines(){
	j := len(index)-1
	for ; j >= 0;{
		defer fmt.Printf("starting word %d ending word %d\n", index[j]+1, j+1)
		j = index[j]-1
	}
}

