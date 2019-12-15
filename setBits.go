package main

import "fmt"

func main(){
  for i := 1; i < 34; i++{
    fmt.Println(i, setBits(i))
  }
}

//Total number of setBits from 1 to N, Time - O(msb_pos_N)
func setBits(N int)int{
  if N == 1{
    return 1
  }
  a := []int{1}
  var i uint
  var res int
  for ; 4<<i < N+2; i++{
    a = append(a, 2*a[len(a)-1]+2<<i)
  }
  //fmt.Println(a) //a = []int{1,4,12,32,80.....}, i = len(a)-1
  for ;N > 0 && i >= 0; i--{
    if N == 2<<i - 1{
      res += a[i]
      break
    }else if N >= 2<<i{
      N -= 2<<i
      res += a[i] + 1 + N
    }
  }
  return res
}
