package main

import "fmt"

//find edges such that removing them makes some verteces unreachable
func main(){
  connections := [][][]int{
    {{0,1}},
    {{0,1},{1,2},{2,0},{1,3}},
    {{0,1},{1,2},{2,0},{0,3},{3,4}},
    {{0,1},{1,2},{2,3}},
    {{0,1},{1,2},{2,0},{1,6},{1,3},{3,4},{4,5},{5,1}},
    {{0,1},{1,2},{2,3},{3,4},{4,5},{5,6},{3,7},{5,7},{5,8},{2,5}},
  }
  // answers
  /*
  [[0,1]]
  [[1,3]]
  [[0,3],[3,4]]
  [[0,1],[1,2],[2,3]]
  [[1,6]]
  [[0,1],[1,2],[5,6],[5,8]]
  */
}

func criticalConnections(n int, connections [][]int) [][]int {
  if len(connections) == n-1{
    return connections
  }
  var i, v int
  var res [][]int
  anc := make([]int, n)
  s := [][]int{{},{0,0}} //s[i]-[]int{vertex, index in a[vertex]}
  vv := make([]int, n) //visited_verteces_index_stack
  a := make([][][]int, n) //adjacency_list
  e  := make([]bool, len(connections)) //edges[i] = []bool{visited}
  for ; i < len(connections); i++{
    a[connections[i][0]] = append(a[connections[i][0]], []int{i, connections[i][1]})
    a[connections[i][1]] = append(a[connections[i][1]], []int{i, connections[i][0]})
  }
  //fmt.Println(a)
  vv[0] = 1
  for ;len(s) > 1;{
    v = s[len(s)-1][0]
    i = s[len(s)-1][1]
    if i > 0{
      if anc[a[v][i-1][1]] > len(s)-1{
        res = append(res, connections[a[v][i-1][0]])
      }else if anc[a[v][i-1][1]] < anc[v]{
        anc[v] = anc[a[v][i-1][1]]
      }
    }

    for ;i < len(a[v]); i++{
      if e[a[v][i][0]]{ //edge already visited
        continue
      }else if vv[a[v][i][1]] > 0{ //destination vertex already visited
        e[a[v][i][0]] = true
        if vv[a[v][i][1]] < anc[v]{
          anc[v] = vv[a[v][i][1]]
        }
      }else{
        e[a[v][i][0]] = true
        break
      }
    }

    if i == len(a[v]){
      s = s[:len(s)-1]
    }else{
      s[len(s)-1][1] = i+1
      vv[a[v][i][1]] = len(s)
      anc[a[v][i][1]] = n+2
      s = append(s, []int{a[v][i][1],0})
    }
  }
  return res
}
