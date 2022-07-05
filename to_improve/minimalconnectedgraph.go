package main

import "fmt"

func main(){
  adjlist := [][]int{{1},{2},{3},{4,6},{5},{6},{3,5}}
  fmt.Println(isminimalconnected(adjlist))
}

func isminimalconnected(adjlist [][]int)bool{
  if len(adjlist) == 1{
    return true
  }
  visited := map[int]bool{0: true}
  /*                                            //vertex number
    visited, acyclic = isacyclic(adjlist, visited, 0)
    if !acyclic || len(visited) < len(adjlist){
      return false
    }
    return false
  */
  for i := 0; i < len(adjlist[0]); i++{
      if !isacyclic(adjlist, visited, adjlist[0][i], 0){
        return false
      }
  }

  if len(visited) < len(adjlist){ //disconnected graph
    return false
  }

  return true
}

func isacyclic(adjlist [][]int, visited map[int]bool, i, parent int)bool{
  if visited[i]{
    return false
  }
  visited[i] = true
  for j := 0; j < len(adjlist[i]); j++{
    if adjlist[i][j] == parent{
      continue
    }else{
      if !isacyclic(adjlist, visited, adjlist[i][j], i){
        return false
      }
    }
  }
  return true
}
