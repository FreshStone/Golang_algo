package main

import "fmt"

func main(){
  s := "racecarannakayak"
  fmt.Println(f(s))
}

func f(s string)[]string{
  if len(s) == 0{
    return []string{}
  }
  type n struct{
    len, start int
  }
  dp := make([]n, len(s)+1)
  dp[0] = n{len: 0, start: -1}
  dp[1] = n{len: 1, start: 0}
  tmp := s[:1]
  var tmp2 string
  var id int //index with min len([]string)
  for i := 1; i < len(s); i++{
    id = i
    tmp2 = tmp
    for l := i+1; l >1; l--{
      if tmp2 == s[i+1-l/2:i+1] && dp[i-l+1].len < dp[id].len{
        id = i-l+1
      }
      tmp2 = tmp2[:len(tmp2)-1]
      if l%2 != 0{
        tmp2 = string(s[i-l/2]) + tmp2
      }
    }
    dp[i+1].len = dp[id].len + 1
    dp[i+1].start = id
    if i%2 == 0{
      tmp = string(s[i/2]) + tmp
    }
  }
  res := make([]string, dp[len(dp)-1].len)
  id = len(dp)-1
  for ;id > 0; {
    res[dp[id].len -1] = s[dp[id].start:id]
    id = dp[id].start
  }
  return res
}
