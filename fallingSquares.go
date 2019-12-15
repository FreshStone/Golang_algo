package main

import (
  "fmt"
)

func main(){
  positions := [][]int{{1,2},{4,3},{3,2},{2,4},{3,4}}
//  positions := [][]int{{9,1},{6,5},{6,7}}
  fmt.Println(fallingSquares(positions))
}

func fallingSquares(positions [][]int) []int {
  if len(positions) == 1{
    return []int{positions[0][1]}
  }
  var i, j, l, h, max int
  res := []int{positions[0][1]}
  dp := [][]int{{positions[0][0], positions[0][0]+positions[0][1], positions[0][1]}}
  for i = 1; i < len(positions); i++{
    l = 0
    h = len(dp)-1
    for ;l <= h;{
      if dp[(l+h)/2][1] <= positions[i][0]{
        l = (l+h)/2 + 1
      }else if dp[(l+h)/2][0] > positions[i][0]{
        h = (l+h)/2 - 1
      }else{
        l = (l+h)/2
        break
      }
    }
    j = l
    h = len(dp)-1
    for ;j <= h;{
      if dp[(j+h)/2][1] < positions[i][0]+positions[i][1]{
        j = (j+h)/2 + 1
      }else if dp[(j+h)/2][0] >= positions[i][0]+positions[i][1]{
        h = (j+h)/2 - 1
      }else{
        h = (j+h)/2
        break
      }
    }
    //fmt.Println(l,h)
    if l > h{
      max = positions[i][1]
      dp = append(dp[:l], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[l:]...)...)
    }else if l == h{
      max = dp[l][2] + positions[i][1]
      if dp[l][0] < positions[i][0]{
        if dp[h][1] > positions[i][0]+positions[i][1]{
          dp = append(dp[:l+1], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max},{0, dp[l][1], dp[l][2]}}, dp[h+1:]...)...)
          dp[l][1] = dp[l+1][0]-1
          dp[l+2][0] = dp[l+1][1]+1
        }else{
          dp[l][1] = positions[i][0]-1
          dp = append(dp[:l+1], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h+1:]...)...)
        }
      }else{
        if dp[h][1] > positions[i][0]+positions[i][1]{
          dp[h][0] = positions[i][0]+positions[i][1]+1
          dp = append(dp[:l], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h:]...)...)
        }else{
          dp = append(dp[:l], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h+1:]...)...)
        }
      }
    }else{
      max = dp[l][2]
      for j = l+1; j <= h; j++{
        if max < dp[j][2]{
          max = dp[j][2]
        }
      }
      max += positions[i][1]
      //fmt.Println(l, dp[l][0], positions[i][0])
      if dp[l][0] < positions[i][0]{
        dp[l][1] = positions[i][0]-1
        if dp[h][1] > positions[i][0]+positions[i][1]{
          dp[h][0] = positions[i][0]+positions[i][1] + 1
          dp = append(dp[:l+1], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h:]...)...)
        }else{
          dp = append(dp[:l+1], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h+1:]...)...)
        }
      }else{
        if dp[h][1] > positions[i][0]+positions[i][1]{
          dp[h][0] = positions[i][0]+positions[i][1] + 1
          dp = append(dp[:l], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h:]...)...)
        }else{
          dp = append(dp[:l], append([][]int{{positions[i][0], positions[i][0]+positions[i][1], max}}, dp[h+1:]...)...)
        }
      }
    }
    fmt.Println(dp, max)
    if max > res[i-1]{
      res = append(res, max)
    }else{
      res = append(res, res[i-1])
    }
  }
  return res
}
