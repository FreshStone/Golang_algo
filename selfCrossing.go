package main

import (
  "fmt"
)

func main(){
  tests := [][]int{
    {2,1,1,2},
    {1,2,3,4},
    {1,1,1,1},
    {1,1,2,2,2,1},
    {1,1,2,2,3,1,1},
    {2,2,3,4,4,1,2,},
    {3,3,4,2,2},
    {3,3,2,1,1},
  }
  ans := []bool{true, false, true, true, true, false, false, false}
  for i, x := range tests{
    if isSelfCrossing(x) != ans[i]{
      fmt.Println("wrong", i, x)
    }else{
      fmt.Println("correct")
    }
  }
}

func isSelfCrossing(x []int) bool {
  if len(x) < 4{
    return false
  }
  var i int
  var compare bool
  v := [][]int{{0, 0, x[0]}, {-x[1], x[0], x[0]-x[2]}} // x, y1, y2
  h := [][]int{{x[0], 0, -x[1]}, {x[0],-x[1],0}} // y, x1, x2
  if x[0] < x[3]{
    compare = true
  }
  for i = 3; i < len(x); i++{
    switch i%4{
    case 0: //north
      if compare{
        if h[1][0] + x[i] < v[0][1]{
          compare = false
        }else if h[1][0] + x[i] > v[0][2]{
          compare = true
        }else{
          v[1][0] = v[0][0]
          v[1][1] = v[0][2]
          v[1][2] = v[0][1]
          compare = false
        }
      }else if h[1][0]+x[i] >= h[0][0]{
        return true
      }
      v[0][0] = h[1][2]
      v[0][1] = h[1][0]
      v[0][2] = h[1][0]+x[i]
    case 1: //west
      if compare{
        if v[0][0] - x[i] > h[0][1]{
          compare = false
        }else if v[0][0] - x[i] < h[0][2]{
          compare = true
        }else{
          h[1][0] = h[0][0]
          h[1][1] = h[0][2]
          h[1][2] = h[0][1]
          compare = false
        }
      }else if v[0][0] - x[i] <= v[1][0]{ //intersecting
        return true
      }
      h[0][0] = v[0][2]
      h[0][1] = v[0][0]
      h[0][2] = v[0][0]-x[i]
    case 2: //south
      if compare{
        if h[0][0]-x[i] > v[1][1]{
          compare = false
        }else if h[0][0] - x[i] < v[1][2]{
          compare = true
        }else{
          v[0][0] = v[1][0]
          v[0][1] = v[1][2]
          v[0][2] = v[1][1]
          compare = false
        }
      }else if h[0][0] - x[i] <= h[1][0]{
        return true
      }
      v[1][0] = h[0][2]
      v[1][1] = h[0][0]
      v[1][2] = h[0][0] - x[i]
    case 3: //east
      if compare{
        if v[1][0] +x[i] < h[1][1]{
          compare = false
        }else if v[1][0] + x[i] > h[1][2]{
          compare = true
        }else{
          h[0][0] = h[1][0]
          h[0][1] = h[1][2]
          h[0][2] = h[1][1]
          compare = false
        }
      }else if v[1][0] + x[i] >= v[0][0]{ //intersecting
        return true
      }
      h[1][0] = v[1][2]
      h[1][1] = v[1][0]
      h[1][2] = v[1][0]+x[i]
    }
  }
  return false
}
