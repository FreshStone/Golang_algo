package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main(){
  s := "055255255255"
  fmt.Println(f(s))
  //fmt.Println(frec(ip,"", 0, 0, 0, []string{}))
}

func f(s string)[]string{
  if len(s) > 12 || len(s) < 4{
    return []string{}
  }
  var res []string
  var tmp strings.Builder
  var j, id, v int
  dp := make([][][]int, 4)
  dp[0] = [][]int{{len(s)-1}}
  for i := len(s)-2; i >= 0; i--{
    j = 3
    if len(s)-1-i < 3{
      j = len(s)-1-i
    }
    for ; j >= 0; j--{
      for id = 0; id < len(dp[j]);{
        if len(dp[j][id]) == 0{
		        break
	      }
        if dp[j][id][0] - i > 2{
	         dp[j] = append(dp[j][:id], dp[j][id+1:]...)
	      }else{
          v,_ = strconv.Atoi(s[i:dp[j][id][0]+1])
          if v > 255{
            dp[j] = append(dp[j][:id], dp[j][id+1:]...)
          }else{
	          id += 1
	        }
        }
      }
      if j > 0{
          for id, _ = range dp[j-1]{
            dp[j] = append(dp[j], append([]int{i}, dp[j-1][id]...))
          }
      }
    }
  }

  for j, _ = range dp[3]{
    tmp.WriteString(s[:dp[3][j][0]+1])
    tmp.WriteString(".")
    tmp.WriteString(s[dp[3][j][0]+1:dp[3][j][1]+1])
    tmp.WriteString(".")
    tmp.WriteString(s[dp[3][j][1]+1:dp[3][j][2]+1])
    tmp.WriteString(".")
    tmp.WriteString(s[dp[3][j][2]+1:])
    res = append(res, tmp.String())
    tmp.Reset()
  }
  return res
}

func frec(ip, tmp string, curr, prev_dot, dots int, ip_ranges []string)[]string{
  if (len(ip) - curr < 3-dots) || (len(ip) - curr > 3*(4-dots)){
    return ip_ranges
  }
  if curr == len(ip){
    if dots == 3{
      tmp += ip[prev_dot:curr]
      ip_ranges = append(ip_ranges, tmp)
    }
    return ip_ranges
  }
  v, _ := strconv.Atoi(ip[prev_dot:curr+1])
  if v < 256{
    ip_ranges = frec(ip, tmp, curr+1, prev_dot, dots, ip_ranges)
  }
  if curr != 0{
    tmp += ip[prev_dot:curr]
    tmp += "."
    return frec(ip, tmp, curr+1, curr, dots+1, ip_ranges)
  }
  return ip_ranges
}
