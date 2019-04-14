package main

import (
	"fmt"
)

func main(){
	n := 161
	fmt.Println(knightMoves(n))
}

func knightMoves(N int) int{
	if N ==1 {
		return 10
	}
	var r int
	max := 1000000007
	dp := make([][4]int, 2)
	m := [][]int{{1,2},{0,0},{0,0,3},{2,2}}
	dp[0][0] = 1
	dp[0][1] = 1
	dp[0][2] = 1
	dp[0][3] = 1
	for i := 1; i<N ; i++{
		r = i%2
		for j := 0; j< len(dp[0]); j++{
			dp[r][j] = 0
			if j == 2{
	                        dp[r][2] = dp[r][1] + dp[1-r][3]
                        }else{
	                        dp[r][j] = dp[1-r][m[j][0]] + dp[1-r][m[j][1]]
			}
			dp[r][j] = dp[r][j]%max//(int(math.Pow(10,9))+7)
		}
	}
	return (4*dp[r][0]+2*dp[r][1]+2*dp[r][2]+dp[r][3])%max
}
