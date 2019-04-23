package main

import (
	"fmt"
	"math"
)

func main(){
	b := 1000
	d := 15
	t := 60
	fmt.Println(poorPigs(b,d,t))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1{
		return 0
	}
	itr := minutesToTest/minutesToDie
	var dp [][]int
	dp = append(dp, make([]int, itr))
	for j := 0; j < itr; j++{
		if j == 0{
			dp[0][0] = 2
		}else{
			dp[0][j] = 1+dp[0][j-1]
		}

		if dp[0][j] >= buckets{
			return len(dp)
		}
	}

	var c, a int
	for {
		dp = append(dp, make([]int, itr))
		for j := 0; j < itr; j++{
			if j == 0{
				dp[len(dp)-1][0] = int(math.Pow(2, float64(len(dp))))
			}else{
				c = len(dp)
				a = 0
				for k := 2; k < len(dp); k++{
					c = (c*(len(dp)-k+1))/(k)
					a = a + c*dp[len(dp)-k-1][j-1]
				}
				dp[len(dp)-1][j] = len(dp)*dp[len(dp)-2][j-1] + a + dp[len(dp)-1][j-1] + 1
				/*
				optimised soln
				if j > 0{
					dp[len(dp)-1][j] = int(math.Pow(float64(1+dp[0][j-1]), float64(len(dp))))
				}
				*/
			}

			if dp[len(dp)-1][j] >= buckets{
				return len(dp)
			}
		}
	}
	return -1
}

