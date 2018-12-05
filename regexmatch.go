package main

import (
	"fmt"
)

func main() {
	p := "a*b.*"
	s := "aabce"
	dp := make([][]bool, len(p)+1)
	for i, _ := range dp{
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true

	for i := 1; i < len(p)+1; i++ {
		for j := 1; j < len(s)+1; j++{
			if (p[i-1] == s[j-1] || p[i-1] == 46){
				dp[i][j] = dp[i-1][j-1]
			}else if p[i-1] == 42{
				if ( p[i-2] == s[j-1] || p[i-2] == 46){
					dp[i][j] = (dp[i-1][j] || dp[i][j-1])
				}else {
					dp[i][j] = dp[i-2][j] //zero occurence of character before *
				}
			}else {
				dp[i][j] = false
			}
		}
	}
fmt.Println(dp[len(p)][len(s)])
}

