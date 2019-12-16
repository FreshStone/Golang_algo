package main

import "fmt"

func main(){
  N := 7
  fmt.Println(numTilings(N))
}

func numTilings(N int) int {
    if N < 3{
        return N
    }
    dp := make([][]int, N+1)
    dp[0] = []int{1,0}
    dp[1] = []int{1,0}
    dp[2] = []int{2,0}
    for i := 3; i < N+1; i++{
        if i > 27{
            dp[i] = []int{0, dp[i-1][1]+dp[i-3][0]%1000000007}
            dp[i][0] = (dp[i-1][0] + dp[i-2][0] + 2*dp[i][1])%1000000007
        }else{
            dp[i] = []int{0, dp[i-1][1]+dp[i-3][0]}
            dp[i][0] = dp[i-1][0] + dp[i-2][0] + 2*dp[i][1]
        }
    }
    return dp[N][0]
}
