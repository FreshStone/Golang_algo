package main

import (
  "fmt"
)

func main(){
  //D := []string{"1","3","5","7"}
  //a := []int{100, 16, 15, 3, 4, 1}
  D := []string{"1","4","9"}
  N := 1000000000
  //for _ , N := range a{
    fmt.Println(atMostNGivenDigitSet(D, N))
  //}
}

func atMostNGivenDigitSet(D []string, N int) int {
  if len(D) == 0{
    return 0
  }
  p := make([]bool, 10) //present in D
  s := make([]int, 10) //smaller/equal to i+1
  var res, r, l, e, d int //d- #DigitsInN
  for r = 1; r < 10; r++{
    if l < len(D) && int(D[l][0] - 48) == r{
      p[r] = true
      s[r] = s[r-1]+1
      l += 1
    }else{
      s[r] = s[r-1]
    }
  }
  l = 1
  for ;N > 0; {
    r = N%10
    if p[r]{
      res += (s[r]-1)*l
      e += 1
    }else{
      res = s[r]*l
    }
    N = N/10
    l *= len(D)
    d += 1
  }
  if e == d{
    res += 1
  }
  if len(D) == 1{
    res += d-1
  }else{
    res += (l-len(D))/(len(D)-1) //G.P sum - (len(D)*(l/len(D)-1))/(len(D)-1)
  }
  return res
}
