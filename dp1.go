/*Different ways to sum n using numbers greater than or equal to m*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("insert n and m")
	}
	a, _ := strconv.ParseInt(os.Args[1], 10, 64)
	b, _ := strconv.ParseInt(os.Args[2], 10, 64)
	n := int(a)
	m := int(b)
	val := make([]int, n-m+1)
	dp := make([][]int, n-m+1)
	for i,_ := range dp{
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
		val[i] = m + i
	}

	for i := 0; i < n-m+1; i++{
		for j := 1; j < n+1; j++{
			if i == 0 {
				/*
				if val[i] > j{
					dp[i][j] = 0
				}else{
					if dp[i][j-val[i]] == 0{
                                        dp[i][j] = 0
                                        }else{
                                        dp[i][j] = 1 
                                        }
				}
				*/
				if (val[i] <= j) && (dp[i][j-val[i]] != 0) {
                                        dp[i][j] = 1
                                }else{
                                        dp[i][j] = 0
                                }
			}else{
				/*
				if val[i] > j {
				dp[i][j] = dp[i-1][j]
				}else{
					if dp[i][j-val[i]] == 0{
					dp[i][j] = dp[i-1][j]
					}else{
					dp[i][j] = 1 + dp[i-1][j]
					}
				}
				*/
				if (val[i] <= j) && (dp[i][j-val[i]] != 0) {
					dp[i][j] = 1 + dp[i-1][j]
				}else{
					dp[i][j] = dp[i-1][j]
				}
			}

	  }
  }
	  fmt.Printf("%v \n", dp[n-m][n])
}

