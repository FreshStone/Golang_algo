package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main(){
	var N int
	var a, b int64
	var fl bool
 var s []string
	edges := [][]int{}
	file, err := os.Open("/home/shady/Downloads/fileInput.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  
 for scanner.Scan() {
 							s = strings.Split(scanner.Text(), " ")
								a, _ = strconv.ParseInt(s[0], 10, 64)
 							if fl{
									b, _ = strconv.ParseInt(s[1], 10, 64)
									edges = append(edges, []int{int(a),int(b)})
								}else{
									fl = true
									N = int(a)
								}
 }

	file.Close()
	//fmt.Println(edges[:100])
	fmt.Println(captainAmerica_rec(N,edges))
}

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
  onz := make(map[int]bool)//out_degree of scc is zero or not

  for ; i < len(edges); i++{
    if len(graph[edges[i][0]]) == 0{
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
  //fmt.Println(onz)
  /*if there are more than 1 SCC with out degree 0 than return 0 
  else 
  return the number of nodes in that SCC (out degree = 0)*/
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
      //fmt.Print((*stack)[i], " ") //prints all nodes of a SCC
      (*stack) = (*stack)[:i]
    }
    //fmt.Println()
  }
  return
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}