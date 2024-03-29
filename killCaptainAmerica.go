package main

import "fmt"

func main(){
  N := []int{5,6,8,7,3,3,3,3,13,13,8}
  testGraphs := [][][]int{
    {{1,2},{2,3},{3,4},{4,3},{5,4}},
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
  ans := []int{2,6,1,3,3,1,0,1,3,0,1}
  for i := 0 ; i < len(testGraphs); i++{
    if ans[i] != captainAmerica_rec(N[i], testGraphs[i]){//captainAmerica(N[i], testGraphs[i]){
      fmt.Println("error", testGraphs[i], "expected", ans[i], "returning", captainAmerica_rec(N[i], testGraphs[i]))
    }else{
      fmt.Println("correct answer")
    }
  }

}
/* 
important learnings--

if we condense each scc component of a directed graph into a node ( that means no cycle exist), 
than the following holds true-
1. there must be atleast one node with its out degree = zero. (to avoid cycles)

2. if a node is reachable by all nodes than its out degree must be zero.(otherwise it will cause cycles)

*/

func captainAmerica_rec(n int, edges [][]int)int{
  //n, len(v) > 1
  if len(edges) < n-1{
    return 0
  }

  if n == 1{
    return 1
  }

  var i, j, y, time int
  var found bool
  visited := make([]int, n+1) //because node numbering starts from 1
  low := make([]int, n+1)
  in_stack := make([]bool, n+1)
  stack := []int{}
  graph := make([][]int, n+1)
  onz := make(map[int]bool)//out_degree of a scc is zero or not (in_stack could also have been used to save memory)

  for ; i < len(edges); i++{
    if len(graph[edges[i][0]]) == 0{
      j -= 1
    }
    graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
  }

  if j > 1{ // checks for terminal nodes
    return 0
  }

  for i = 1; i < n+1; i++{
    if visited[i] ==  0{
      rec(i, &time, visited, low, &stack, in_stack, graph)
    }
  }
  //fmt.Println(low)

  /*labelling SCC's with out degree != 0 */
  for i =0 ; i < len(edges); i++{
    if low[edges[i][0]] != low[edges[i][1]]{
      onz[low[edges[i][0]]] = true
    }
  }
  /*if there are more than 1 SCC with out degree 0 than return 0 
  else 
  return the number of nodes in that SCC (with out degree = 0)*/
  for i, j = 1, 0; i < n+1; i++{
    if onz[low[i]] == false{ //out degree of scc is zero 
      if found{
        if low[i] == y{
          j += 1
        }else{
          return 0
        }
      }else{
        found = true
        y = low[i]
        j = 1
      }
    }
  }
  return j
}

func rec(node int, time *int, visited []int, low []int, stack *[]int, in_stack []bool, graph [][]int){
  var i, j int
  *time += 1
  low[node] = *time
  visited[node] = *time
  *stack = append(*stack, node)
  in_stack[node] = true

  for i = 0; i < len(graph[node]); i++{
    j = graph[node][i]
    if visited[j] == 0{
      rec(j,time,visited,low,stack,in_stack,graph)
      low[node] = min(low[node], low[j])
    }else if in_stack[j]{ //back edge
      low[node] = min(low[node], low[j])
    }
  }

  if visited[node] == low[node]{
    for i = len(*stack)-1; i >= 0 && low[(*stack)[i]] >= low[node]; i--{
      low[(*stack)[i]] = low[node]
      in_stack[(*stack)[i]] = false
     // fmt.Print((*stack)[i], " ") //prints all nodes of a SCC
      (*stack) = (*stack)[:i]
    }
   // fmt.Println()
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
  var found bool
  var time int = 1
  visited := make([]int, n+1) //because node numbering starts from 1
  low := make([]int, n+1)
  in_stack := make([]bool, n+1)
  rec_stack := [][]int{} //rec_stack[i] = []int{node, neighbor_to_visit}
  stack := []int{}
  graph := make([][]int, n+1)
  onz := make(map[int]bool) //low_id -> strongly connected components

  for ; i < len(edges); i++{
    if len(graph[edges[i][0]]) == 0{
      j -= 1
    }
    graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
  }

  if j > 1{ //nodes with out degree = 0
    return 0
  }

  for i = 1; i < n+1; i++{
    if visited[i] == 0{
      rec_stack = append(rec_stack, []int{i,0})
      stack = append(stack, i)
      for len(rec_stack) > 0 {
        x = rec_stack[len(rec_stack)-1][0]
        y = rec_stack[len(rec_stack)-1][1]
        if visited[x] == 0{
          visited[x] = time
          low[x] = time
          in_stack[x] = true
          time++
        }else if in_stack[graph[x][y-1]] {
          low[x] = min(low[x], low[graph[x][y-1]])
        }
        //fmt.Println(x, y, low[x])
        if len(graph[x]) == y{
          rec_stack = rec_stack[:len(rec_stack)-1]
          if low[x] == visited[x]{
            for j = len(stack)-1; j >= 0 && low[stack[j]] >= low[x]; j--{
              in_stack[stack[j]] = false // finds all nodes 
              low[stack[j]] = low[x]    // of a SCC
            }
            stack = stack[:j+1]
          }
        }else if visited[graph[x][y]] == 0{
          rec_stack = append(rec_stack, []int{graph[x][y],0})
          stack = append(stack, graph[x][y])
          rec_stack[len(rec_stack)-2][1]++
        }else if in_stack[graph[x][y]] { //back edge
          low[x] = min(low[x], low[graph[x][y]])
          rec_stack[len(rec_stack)-1][1]++
        }else{ //cross edge (pointing to node outsode scc)
          rec_stack[len(rec_stack)-1][1]++
        }
      }
    }
  }
  //fmt.Println(low)
  /*labelling SCC's with out degree = 0*/
  for i =0 ; i < len(edges); i++{
    if low[edges[i][0]] != low[edges[i][1]]{
      onz[low[edges[i][0]]] = true
    }
  }
  /*if there are more than 1 SCC with out degree 0 than return 0 
  else 
  return the number of nodes in that SCC*/
  for i, j = 1, 0; i < n+1; i++{
    if onz[low[i]] == false{ //out degree of scc is zero 
      if found{
        if low[i] == y{
          j += 1
        }else{
          return 0
        }
      }else{
        found = true
        y = low[i]
        j = 1
      }
    }
  }
  return j
}

func min(a, b int)int{
  if a < b{
    return a
  }
  return b
}
