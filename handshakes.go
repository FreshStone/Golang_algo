package main

import "fmt"

func main(){
  tests := []int{100}
  for _, num_people := range tests{
    fmt.Println(numberOfWays(num_people))
  }
}

func numberOfWays(num_people int)int{
  if num_people == 0{
    return 0
  }
  var i, j int
  dp := make([]int, num_people+1)
  dp[0] = 1
  for i = 2; i <= num_people; {
    for j = 2; j <= i; {
      dp[i] = (dp[i] + dp[j-2]*dp[i-j])%1000000007
      j += 2
    }
    i += 2
  }
  //fmt.Println(dp)
  return dp[num_people]
}
