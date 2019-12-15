package main

import (
  "fmt"
  "sort"
)

type edge struct{
  src, des, weight int
}
type vertex struct{
  par, rank int
}
func main(){
  g := [][][]int{
    {{1,4},{7,8}},
    {{2, 8},{7,11}},
    {{3,7},{5,4},{8,2}},
    {{4,9},{5,14}},
    {{5,10}},
    {{6,2}},
    {{7,1},{8,6}},
    {{8,7}},
    {},
  }
  fmt.Println(mst(g))
}

func mst(g [][][]int)[][]int{
  var arr []edge
  var res [][]int
  s := make([]vertex, len(g))
  var i, j, t, r1, r2, e int
  for ;i < len(g); i++{
    s[i] = vertex{par: i}
    for j = 0; j < len(g[i]); j++{
      arr = append(arr, edge{src:i, des: g[i][j][0], weight: g[i][j][1]})
    }
  }
  sort.Slice(arr, func(i int, j int) bool {
    return arr[i].weight < arr[j].weight
  })
  for i = 0; e < len(g)-1 && i < len(arr); i++{         //{
    j = findparent(s, arr[i].src)
    t = findparent(s, arr[i].des)
    if j != t{
      e += 1
      res = append(res, []int{arr[i].src, arr[i].des})
      r1 = s[j].rank //rank of parent                       Union of subset of vertices
      r2 = s[t].rank
      if r1 > r2{
        s[t].par = j
      }else if r2 < r1{
        s[j].par = t
      }else{
        s[t].par = j
        s[j].rank += 1
      }
    }
  }                                                       //}
  return res
}

func findparent(s []vertex, i int)int{
  if s[i].par != i{
    s[i].par = findparent(s, s[i].par)
  }
  return s[i].par
}
