package main

import "fmt"

func main(){
	g := []int{2,3,4,9,1,6}
	fmt.Println(OptimalStrategy(g))
}

func OptimalStrategy(g []int)int{
	type score struct{
		winner, looser int
	}

	dp := make([][]score, len(g)+1)
	for i, _ := range dp{
		dp[i] = make([]score, len(g)+1)
	}
	
	for l := 1; l <= len(g); l++{
		i := 0
		for j := l-1; i < len(g) && j < len(g);{
			var s score
			if (g[j] + dp[i][j].looser) > (g[i] + dp[i+1][j+1].looser){
				s.winner = g[j] + dp[i][j].looser
				s.looser = dp[i][j].winner
			}else{
				s.winner = g[i] + dp[i+1][j+1].looser
				s.looser = dp[i+1][j+1].winner
			}
			dp[i][j+1] = s
			i = i+1; j = i-1+l
		}
	}
	return dp[0][len(g)].winner
}

