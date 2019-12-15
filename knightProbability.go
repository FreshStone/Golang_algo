package main

import (
  "fmt"
)

func main(){
    tests := [][]int{
     //N, K, r, c
      {3,2,0,1},
      {3,3,1,1},
    }
    for _, v := range tests{
        fmt.Println(knightProbability(v[0],v[1],v[2], v[3]))
    }
}
/*
  --------------
  |   1  |  2  |
  |      |     |
  --------------
  |   3  |  4  |
  |      |     |
  --------------
  neightbours & prob for a square in Area 1 is same for any of its symmetric square in
  Area 2, 3 & 4
*/
func knightProbability(N int, K int, r int, c int) float64 {
    if K == 0{
      return 1
    }
    if N < 3{
      return 0
    }
    var i, j int
    rch := make([][][]int, (N+1)/2) //reachable squares
    prob := make([][][]float64, len(rch))
    for ;i < len(rch); i++{  //filling each square with nearby squares co-ordinates
      rch[i] = make([][]int, len(rch))
      prob[i] = make([][]float64, len(prob))
      for j = 0; j < len(rch); j++{
        rch[i][j] = []int{}
        prob[i][j] = make([]float64, K)
        if i > 1 && j > 0{
          rch[i][j] = append(rch[i][j], i-2)
          rch[i][j] = append(rch[i][j], j-1)
        }
        if i > 1{// && j < N-1{
          rch[i][j] = append(rch[i][j], i-2)
          if j > len(rch)-2{
            rch[i][j] = append(rch[i][j], N-j-2)
          }else{
            rch[i][j] = append(rch[i][j], j+1)
          }
        }
        if i > 0 && j > 1{
          rch[i][j] = append(rch[i][j], i-1)
          rch[i][j] = append(rch[i][j], j-2)
        }
        if i > 0 && j < N-2{
          rch[i][j] = append(rch[i][j], i-1)
          if j > len(rch)-3{
            rch[i][j] = append(rch[i][j], N-j-3)
          }else{
            rch[i][j] = append(rch[i][j], j+2)
          }
        }
        if j > 1{//&& i < N-1
          if i > len(rch)-2{
            rch[i][j] = append(rch[i][j], N-i-2)
          }else{
            rch[i][j] = append(rch[i][j], i+1)
          }
          rch[i][j] = append(rch[i][j], j-2)
        }
        if j < N-2{ //&& i < N-1
          if i > len(rch)-2{
            rch[i][j] = append(rch[i][j], N-i-2)
          }else{
            rch[i][j] = append(rch[i][j], i+1)
          }
          if j > len(rch)-3{
            rch[i][j] = append(rch[i][j], N-j-3)
          }else{
            rch[i][j] = append(rch[i][j], j+2)
          }
        }
        if i < N-2 && j > 0{
          if i > len(rch)-3{
            rch[i][j] = append(rch[i][j], N-i-3)
          }else{
            rch[i][j] = append(rch[i][j], i+2)
          }
          rch[i][j] = append(rch[i][j], j-1)
        }
        if i < N-2{// && j < N-1{
          if i > len(rch)-3{
            rch[i][j] = append(rch[i][j], N-i-3)
          }else{
            rch[i][j] = append(rch[i][j], i+2)
          }
          if j > len(rch)-2{
            rch[i][j] = append(rch[i][j], N-j-2)
          }else{
            rch[i][j] = append(rch[i][j], j+1)
          }
        }
        prob[i][j][0] = float64(len(rch[i][j])/2)/8.0
      }
    }
    //fmt.Println(prob)
    //fmt.Println(rch)
    if r > len(rch)-1{
      r = N-r-1
    }
    if c > len(rch)-1{
      c = N-c-1
    }
    rec(r, c, K, rch, prob)
    return prob[r][c][K-1]
}

func rec(r, c, k int, rch [][][]int, prob [][][]float64){
  if prob[r][c][k-1] > 0{
    return
  }
  for i := 0; i < len(rch[r][c]);{
    rec(rch[r][c][i], rch[r][c][i+1], k-1, rch, prob)
    prob[r][c][k-1] += prob[rch[r][c][i]][rch[r][c][i+1]][k-2]
    i += 2
  }
  prob[r][c][k-1] /= 8
  return
}
