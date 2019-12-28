package main

import (
  "fmt"
)

func main(){
  //n, l, k, ans
  tests := [][]int{
    {3, 3, 1, 6},
    {2, 3, 0, 6},
    {2, 3, 1, 2},
    {4, 6, 2, 168},
    {4, 5, 2, 72},
    {8, 15, 2, 835297215},
    {8, 15, 3, 617407615},
    {8, 15, 4, 876639965},
    {8, 15, 5, 376185600},
    {8, 15, 6, 10281600},
    {8, 15, 7, 40320},
    {10, 15, 6, 195775804},
    {7, 12, 4, 4868640},
    {25,28,5,906368459},
    {25,26,5,553800836},
  }
  for i, _ := range tests{
    if numMusicPlaylists(tests[i][0], tests[i][1], tests[i][2]) != tests[i][3]{
      fmt.Println("test failed", i)
      break
    }
  }
}

func numMusicPlaylists(N int, L int, K int)int{
  if N == L{
    return factorial(N)%1000000007
  }
  dp := make([]int, L-N+1)
  dp[0] = 1
  var i, j int
  for ; i < N-K; i++{
    for j = 1; j < len(dp); j++{
      dp[j] += (dp[j-1]*(N-i-K)+dp[j])%1000000007
    }
  //  fmt.Println(dp)
  }
  return (factorial(N)*dp[L-N])%1000000007
}

func factorial(n int)int{
  res := 1
  for ; n > 1; n--{
    res = (res*n)%1000000007
  }
  return res
}
