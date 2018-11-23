package main

import (
	"fmt"
	"math"
)
var a = []int{6,1,3,1} //len of consecutive words
var index = make([]int, len(a))
var w = 8 //max width of each line
//var top = -1

func main(){
	dp := make([]int, len(a)+1)
	var n, min_cost int
	dp[0] = 0;
	for i := 1; i <= len(a);i++{
		n = a[i-1]
		min_cost = dp[i-1] + int(math.Pow(float64(w-n), 3))
		j := i-1
		Loop :
			for ; j >=1; j-- {
				n = n + a[j-1] +1

				if n > w  {
					break Loop
				}
				if min_cost >  dp[j-1] + int(math.Pow(float64(w-n), 3)){
					min_cost = dp[j-1] + int(math.Pow(float64(w-n), 3))
					index[i-1] = j// 1st word 0 index
				}
			}
		dp[i] = min_cost
	}
	fmt.Println("min_cost-", dp[len(a)])
	fmt.Println("index array-", index)
	Printlines()
}

func Printlines(){ //optimize
	var s []int
	line_num := 1
	for i := len(index)-1; i>=0;{
		if index[i] == 0{
			index[i] = i+1
		}
		s = append(s, index[i])
		i = index[i]-2
	}
	
	for j := len(s)-1; j >= 0; j--{
		if j == 0{
			fmt.Printf("line-num %d starting word %d ending word %d\n", line_num, s[j], len(a))
		}else{
			fmt.Printf("line-num %d starting word %d ending word %d\n", line_num, s[j], (s[j-1] - 1))
		}
		line_num += 1
	}
}


