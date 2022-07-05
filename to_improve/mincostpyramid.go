package main

import (
  "fmt"
)
/*
We have N (N > 0) stones of various heights laid out in a row. Task is to make a pyramid from given
array of stones. In a pyramid, height of the stones start from 1, increase by 1, until it reaches
some value x, then decreases by 1 until it reaches 1 again i.e. the stones should be
1, 2, 3, 4…x – 1, x, x – 1, x – 2 … 1. All other stones not part of the pyramid should have a
height 0. We cannot move any of the stones from their current position, however, by paying a fee
of 1, we can reduce the heights of the stones. Minimize the cost of building a pyramid.
*/
func main(){
  nums := [][]int{
    {1,1,3,3,2,1},
    {1},
    {1,1},
    {1,1,1},
    {1,1,1,1},
    {1,2},
    {2,1},
    {1, 2, 3, 4, 2, 1},
    {1,5,2},
    {1,2,1},
    {2,4,5,7,8,10},
    {7,5,4,3,1},
    {3,5,6,7,8,10,3,4,5,1,4,5},

  }
  for _, v := range nums{
    fmt.Println(f(v))
  }
}

func f(nums []int)int{
  if len(nums) == 1{
    return nums[0]-1
  }
  res := 1<<31 - 1
  var i, j, tmp int
  s := nums[0]+nums[len(nums)-1]
  l := make([][]int, len(nums)) //l[i][0] - max_width_left, l[i][1] - max_width_right
  //height_pyramid_ith_index = min(max_width_left, max_width_right)+1
  l[0] = []int{0, 0}
  l[len(nums)-1] = []int{0, 0}
  if len(nums) & 1 == 1{
    s += nums[len(nums)/2]
    l[len(nums)/2] = []int{0, 0}
  }
  for i = 1; i < len(nums)-1; i++{
    j = len(nums)-i-1
    if i < len(nums)/2{
      s += nums[i]+nums[j]
      l[i] = []int{0,0}
      l[j] = []int{0,0}
    }
    l[i][0] = l[i-1][0]+1
    if nums[i]-1 < l[i][0]{
      l[i][0] = nums[i]-1
    }
    l[j][1] = l[j+1][1]+1
    if nums[j]-1 < l[j][1]{
      l[j][1] = nums[j]-1
    }
    if i >= len(nums)/2{
      if l[i][0] < l[i][1]{
        tmp = s-(l[i][0]+1)*(l[i][0]+1)
      }else{
        tmp = s-(l[i][1]+1)*(l[i][1]+1)
      }
      if tmp < res{
        res = tmp
      }
      if l[j][0] < l[j][1]{
        tmp = s-(l[j][0]+1)*(l[j][0]+1)
      }else{
        tmp = s-(l[j][1]+1)*(l[j][1]+1)
      }
      if tmp < res{
        res = tmp
      }
    }
  }
  if s-1 < res{
    res = s-1
  }
//  fmt.Println(l)
  return res
}
