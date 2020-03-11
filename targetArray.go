package main

import "fmt"

func main(){
  tests := [][]int{
    {1,1,1,2},
    {9,3,5},
    {8,5},
    {1,21},
  }
  results := []bool{false, true, true, true}
  for i, target := range tests{
    if results[i] == isPossible(target){
      fmt.Println("correct")
    }else{
      fmt.Println(target, "incorrect")
    }
  }
}

func isPossible(target []int) bool {
  if len(target) == 1 && target[0] != 1{
    return false
  }
  h := make([]int, len(target))
  var s, i, j int
  for ;i < len(target); i++{
    s += target[i]
    h[i] = target[i]
    j = i
    for{
      if (j-1)/2 >= 0 && h[(j-1)/2] < h[j]{
        h[j], h[(j-1)/2] = h[(j-1)/2], h[j]
        j = (j-1)/2
      }else{
        break
      }
    }
  }
  //fmt.Println(h, s)
  for ;h[0] > 1; {
    if 2*h[0] > s {
      if s-h[0] == 1{
        return true
      }else if h[0]%(s-h[0]) != 0 && (h[0]%(s-h[0]) < s-h[0]){
        j = h[0]
        //h[0] = 2*h[0] - s
        h[0] = h[0]%(s-h[0])
        s += h[0] - j
        i = 0
        for{
          if 2*i > len(h)-3{
            if 2*i < len(h) -1 && h[i] < h[2*i+1]{
              h[i], h[2*i+1] = h[2*i+1], h[i]
            }
            break
          }else if h[i] >= h[2*i+1]{
            if h[i] < h[2*i+2]{
              h[i], h[2*i+2] = h[2*i+2], h[i]
              i = 2*i+2
            }else{
              break
            }
          }else{
            if h[i] < h[2*i+2] && h[2*i+2] > h[2*i+1]{
              h[i], h[2*i+2] = h[2*i+2], h[i]
              i = 2*i+2
            }else{
              h[i], h[2*i+1] = h[2*i+1], h[i]
              i = 2*i+1
            }
          }
        }
      }else{
        return false
      }
    }else{
    //  fmt.Println(h)
      return false
    }
  }
  //fmt.Println(h)
  return true
}
