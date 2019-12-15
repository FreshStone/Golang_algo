package main

import "fmt"

func main(){
  poured := 1000000000
  r := 99
  c := 99
  fmt.Println(champagneTower(poured, r, c))
}

func champagneTower(poured int, query_row int, query_glass int)float64{
  if poured == 0{
    return 0
  }
  if query_row == 0 {
    return 1
  }
  var i, x, y int
  visited := make([][]bool, query_row+1)
  excess := make([][]float64, query_row+1)
  stack := [][]int{{query_row, query_glass}}
  for ;i < len(excess); i++{
    visited[i] = make([]bool, i+1)
    excess[i] = make([]float64, i+1)
  }
  excess[0][0] = float64(poured-1)
  visited[0][0] = true
  for ;len(stack) > 0;{
    x = stack[len(stack)-1][0]
    y = stack[len(stack)-1][1]
    if y == 0{
      if visited[x-1][y]{
        if excess[x-1][y] >= 2{
          excess[x][y] = excess[x-1][y]/2 - 1
        }
        visited[x][y] = true
      }else{
        stack = append(stack, []int{x-1, y})
        continue
      }
    }else if y == x{
      if visited[x-1][y-1]{
        if excess[x-1][y-1] >= 2{
          excess[x][y] = excess[x-1][y-1]/2 - 1
        }
        visited[x][y] = true
      }else{
        stack = append(stack, []int{x-1, y-1})
        continue
      }
    }else{
      if !visited[x-1][y-1]{
        stack = append(stack, []int{x-1, y-1})
        continue
      }else if !visited[x-1][y]{
        stack = append(stack, []int{x-1, y})
        continue
      }else{
        if excess[x-1][y] + excess[x-1][y-1] >= 2{
          excess[x][y] = (excess[x-1][y]+excess[x-1][y-1])/2 - 1
        }
        visited[x][y] = true
      }
    }
    stack = stack[:len(stack)-1]
  }
//  printarr(excess)
  if excess[x][y] > 0{
    return 1
  }else if query_glass == 0{
    return excess[x-1][y]/2
  }else if query_glass == query_row{
    return excess[x-1][y-1]/2
  }
  return (excess[x-1][y]+excess[x-1][y-1])/2
}

func printarr(a [][]float64){
  for i := 0; i < len(a); i++{
    fmt.Println(a[i])
  }
  return
}
