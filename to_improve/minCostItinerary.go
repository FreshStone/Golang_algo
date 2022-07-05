package main

import "fmt"

func main(){
  edges := [][]string{
    {"JFK", "ATL"},
    {"ATL", "SFO"},
    {"ORD", "LAX"},
    {"LAX", "DFW"},
    {"JFK", "HKG"},
    {"ATL", "ORD"},
    {"JFK", "LAX"},
    {"ATL", "LAX"},
  }
  price := []int{150,400,200,80,800,90,500,150}
  k := 3
  source := "JFK"
  destination := "LAX"
  fmt.Println(mincostitinerary(source, destination, edges, price, k))
}

func mincostitinerary(source, destination string, edges [][]string, price []int, k int)(int, []string){
  if len(edges) == 0{
    return 0, []string{}
  }
  var ok1, ok2 bool
  var i, j, l, s, e, itinerary_len int
  var mc []int //stores maximum cost to reach a destination
  stations := make(map[string]int) //string to int
  station_map := make(map[int]string) //int to string
  var adj_list [][][]int //adj_list[i][j] = []int{destination_index, cost}
  for ;i < len(edges); i++{
    _, ok1 = stations[edges[i][0]]
    _, ok2 = stations[edges[i][1]]
    if !ok1{
      stations[edges[i][0]] = len(adj_list)
      station_map[len(adj_list)] = edges[i][0]
      adj_list = append(adj_list, [][]int{})
      mc = append(mc, 1<<31-1)
    }
    if !ok2{
      stations[edges[i][1]] = len(adj_list)
      station_map[len(adj_list)] = edges[i][1]
      adj_list = append(adj_list, [][]int{})
      mc = append(mc, 1<<31-1)
    }
    adj_list[stations[edges[i][0]]] = append(adj_list[stations[edges[i][0]]], []int{stations[edges[i][1]], price[i]})
  }
  //station code, par index in bfs, initial cost
  bfs := [][]int{{stations[source],-1,0}}
  mc[stations[source]] = 0
  for j = 0; j <= k; j++{
    e = len(bfs)
    for ;s < e; s++{
      if bfs[s][2] == mc[bfs[s][0]]{ //compare if min cost
        if bfs[s][0] == stations[destination]{
          i = s
          itinerary_len = j
        }
        for l = 0; j < k && l < len(adj_list[bfs[s][0]]); l++{
          if bfs[s][2] + +adj_list[bfs[s][0]][l][1] < mc[adj_list[bfs[s][0]][l][0]]{ //add if min
            mc[adj_list[bfs[s][0]][l][0]] = bfs[s][2] + +adj_list[bfs[s][0]][l][1]
            bfs = append(bfs, []int{adj_list[bfs[s][0]][l][0],s,mc[adj_list[bfs[s][0]][l][0]]})
          }
        }
      }
    }
  }
  //fmt.Println(bfs)
  j = itinerary_len
  res := make([]string, j+1)
  for ; j >= 0; j--{
    res[j] = station_map[bfs[i][0]]
    i = bfs[i][1]
  }
  return mc[stations[destination]], res
}
