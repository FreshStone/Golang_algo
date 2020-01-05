package main

import "fmt"

func main(){
  /*prefer := [][]int{
    {8,6,7,5,9},
    {9,7,6,5,8},
    {6,9,5,8,7},
    {9,6,8,7,5},
    {8,5,6,7,9},
    {3,1,4,2,0},
    {1,0,3,2,4},
    {0,2,4,3,1},
    {3,0,2,1,4},
    {1,4,0,2,3},
  }*/
  prefer := [][]int{
    {4,5,3},
    {5,3,4},
    {3,5,4},
    {0,1,2},
    {2,1,0},
    {2,0,1},
  }
  fmt.Println(isStable(prefer))//women - men mapping
}
func changepartner(prefer [][]int, w, m1, m int)bool{
  for j := 0; j < len(prefer[0]); j++{
    if prefer[w][j] == m1{
      return true
    }else if prefer[w][j] == m{
      return false
    }
  }
  return false
}

func isStable(prefer [][]int)map[int]int{
  if len(prefer) == 0{
    return map[int]int{}
  }
  if len(prefer) == 1{
    return map[int]int{0: 0}
  }
  var i, j int
  freemen := make([]int, len(prefer[0]))
  womenpartner := make(map[int]int)
  for ; i < len(freemen); i++{
    freemen[i] = i
    womenpartner[len(freemen)+i] = -1
  }
  for ; len(freemen) > 0;{
    i = freemen[0]
    for j = 0; j < len(prefer[0]); j++{
      if womenpartner[prefer[i][j]] == -1{ //women not taken
        womenpartner[prefer[i][j]] = i
        freemen = freemen[1:]
        break
      }else if changepartner(prefer, prefer[i][j], i, womenpartner[prefer[i][j]]){
        freemen = append(freemen, womenpartner[prefer[i][j]])
        womenpartner[prefer[i][j]] = i
        freemen = freemen[1:]
        break
      }
    }
  }
  return womenpartner
}
