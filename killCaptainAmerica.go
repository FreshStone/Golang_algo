package main

import "fmt"

func main(){
  N := []int{6,8,7,3,3,3,3,13,13,8}
  testGraphs := [][][]int{
    {{1,2},{2,3},{3,4},{3,5},{5,6},{6,4},{4,2},{5,1}},
    {{1,2},{2,3},{3,4},{4,5},{5,6},{6,7},{7,4},{5,8}},
    {{1,2},{2,3},{3,4},{4,2},{3,5},{5,6},{6,7},{7,5}},
    {{1,2},{2,3},{3,1}},
    {{1,2},{1,3},{3,1}},
    {{1,2},{1,3}},
    {{1,2},{2,3}},
    {{1,2},{2,3},{3,4},{4,5},{5,3},{4,9},{9,6},{6,7},{7,8},{8,6},{9,10},{10,11},{11,12},{12,10},{11,13},{13,7}},
    {{1,2},{2,3},{3,4},{4,5},{5,3},{4,9},{9,10},{10,11},{11,12},{12,10},{11,13},{9,6},{6,7},{7,8},{8,6}},
    {{1,2},{2,3},{3,4},{3,5},{5,6},{5,7},{7,8},{8,6},{6,4}},
  }
  ans := []int{6,1,3,3,1,0,1,3,0,1}
  for i := 0 ; i < len(ans); i++{
    if ans[i] != captainAmerica(N[i], testGraphs[i]){//captainAmerica_rec(N[i], testGraphs[i]){
      fmt.Println("error", N[i], testGraphs[i])
    }else{
      fmt.Println("correct answer")
    }
  }

}

func captainAmerica_rec(n int, edges [][]int)int{
  //n, len(v) > 1
  if len(edges) < n-1{
    return 0
  }

  if n == 1{
    return 1
  }

  var i, j, x, y int
  var time int = 1
  var out_zero, found bool
  visited := make([]int, n+1) //because node numbering starts from 1
  low := make([]int, n+1)
  in_stack := make([]bool, n+1)
  stack := []int{-1}
  graph := make([][]int, n+1)
  m := make(map[int][]int)

  for ; i < len(edges); i++{
    if len(graph[edges[i][0]]) == 0{
      graph[edges[i][0]] = []int{}
      j -= 1
    }
    graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
  }

  //fmt.Println(graph)
  if j > 1{
    return 0
  }

  for i = 1; i < n+1; i++{
    if visited[i] ==  0{
      visited[i] = time
      low[i] = time
      stack = append(stack, i)
      in_stack[i] = true
      rec(i, &time, visited, low, &stack, in_stack, graph)
    }
  }

  //fmt.Println(low)
  for i = 1; i < n+1; i++{
    if _, found = m[low[i]]; found{
      m[low[i]] = append(m[low[i]], i)
    }else{
      m[low[i]] = []int{i}
    }
  }
  //fmt.Println(m)
  found = false
  out_zero = true
  for i = range m{
    for j = 0; j < len(m[i]); j++{
      for x = 0; x < len(graph[m[i][j]]); x++{
        if low[graph[m[i][j]][x]] != i{
          out_zero = false
          break
        }
      }
      if !out_zero{
        break
      }
    }

    if out_zero{
      if found{
        return 0
      }
      found = true
      y = len(m[i])
    }
    out_zero = true
  }

  return y
}

func rec(node int, time *int, visited []int, low []int, stack *[]int, in_stack []bool, graph [][]int){
  var i, j int
  for i = 0; i < len(graph[node]); i++{
    j = graph[node][i]
    if visited[j] == 0{
      *time += 1
      low[j] = *time
      visited[j] = *time
      *stack = append(*stack, j)
      in_stack[j] = true
      rec(j,time,visited,low,stack,in_stack,graph)
      low[node] = min(low[node], low[j])
    }else if in_stack[j]{ //back edge
      low[node] = min(low[node], low[j])
    }
  }

  if visited[node] == low[node]{
    for i = len(*stack)-1; i > 0 && low[(*stack)[i]] >= low[node]; i--{
      low[(*stack)[i]] = low[node]
      in_stack[(*stack)[i]] = false
    }
    (*stack) = (*stack)[:i+1]
  }
  return
}

func captainAmerica(n int, edges [][]int)int{
  //n, len(v) > 1
  if len(edges) < n-1{
    return 0
  }

  if n == 1{
    return 1
  }

  var i, j, x, y int
  var out_zero, found bool
  var time int = 1
  visited := make([]int, n+1) //because node numbering starts from 1
  low := make([]int, n+1)
  in_stack := make([]bool, n+1)
  rec_stack := [][]int{{}} //stack[i] = []int{node, neighbor_to_visit}
  stack := []int{-1}
  graph := make([][]int, n+1)
  m := make(map[int][]int) //low_id -> strongly connected components

  for ; i < len(edges); i++{
    if len(graph[edges[i][0]]) == 0{
      graph[edges[i][0]] = []int{}
      j -= 1
    }
    graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
  }

  //fmt.Println(graph)
  if j > 1{
    return 0
  }

  for i = 1; i < n+1; i++{
    if visited[i] == 0{
      visited[i] = time
      low[i] = time
      in_stack[i] = true
      time++
      rec_stack = append(rec_stack, []int{i,0})
      stack = append(stack, i)

      for len(rec_stack) > 1 {
        x = rec_stack[len(rec_stack)-1][0]
        y = rec_stack[len(rec_stack)-1][1]
        if len(graph[x]) == y{
          //remove node
          if len(rec_stack) > 2{
            low[rec_stack[len(rec_stack)-2][0]] = min(low[rec_stack[len(rec_stack)-2][0]], low[x])
          }
          rec_stack = rec_stack[:len(rec_stack)-1]
          if low[x] == visited[x]{
            for j = len(stack)-1; j > 0 && low[stack[j]] >= low[x]; j--{
              in_stack[stack[j]] = false
              low[stack[j]] = low[x]
            }
            stack = stack[:j+1]
          }
        }else if visited[graph[x][y]] == 0{
          visited[graph[x][y]] = time
          low[graph[x][y]] = time
          in_stack[graph[x][y]] = true
          time++
          rec_stack = append(rec_stack, []int{graph[x][y],0})
          stack = append(stack, graph[x][y])
          rec_stack[len(rec_stack)-2][1]++
        }else if in_stack[graph[x][y]] {
          low[x] = min(low[x], low[graph[x][y]])
          rec_stack[len(rec_stack)-1][1]++
        }else{
          rec_stack[len(rec_stack)-1][1]++
        }
      }
    }
  }
  //fmt.Println(low)
  for i = 1; i < n+1; i++{
    if _, found = m[low[i]]; found{
      m[low[i]] = append(m[low[i]], i)
    }else{
      m[low[i]] = []int{i}
    }
  }

  found = false
  out_zero = true
  //fmt.Println(m)
  for i = range m{
    for j = 0; j < len(m[i]); j++{
      for x = 0; x < len(graph[m[i][j]]); x++{
      //  fmt.Println("edge", m[i][j], "->", graph[m[i][j]][x])
        if low[graph[m[i][j]][x]] != i{
          out_zero = false
          break
        }
      }
      if !out_zero{
        break
      }
    }

    if out_zero{
      if found{
        return 0
      }
      found = true
      y = len(m[i])
    }
    out_zero = true
  }

  return y
}

func min(a, b int)int{
  if a < b{
    return a
  }
  return b
}
