package main

import "fmt"

func main(){
	routes := [][][]int{
		{{1,2,7},{3,6,7}},
		{{7,12},{4,5,15},{6},{15,19},{9,12,13}},
		{{2,3},{4,5}},
		{{1,2,3},{4,5}},
		{{1,2,3},{4,5,6}},
	}
	sources := []int{1,15,1,1,1}
	targets := []int{6,12,6,6,6}
	ans := []int{2,-1,-1,-1,-1}
	for i := 0; i < len(routes); i++{
		if ans[i] != numBusesToDestination(routes[i], sources[i], targets[i]){
			fmt.Println("error","expected-", ans[i], "returning-", numBusesToDestination(routes[i], sources[i], targets[i]))
		}else{
			fmt.Println("correct answer")
		}
	}
}

func numBusesToDestination(routes [][]int, source int, target int) int {
    var i, j, l, b, s int
    var ok bool
    bfs := []int{}
    buses := make([]map[int]bool, len(routes))   
    left_buses := make(map[int]bool) //buses which are not not taken yet
    stops := make(map[int]bool) //bus_stops not taken yet

    for ; i < len(routes); i++{
    	buses[i] = map[int]bool{}
    	left_buses[i] = true
    	for j = 0; j < len(routes[i]); j++{
    		buses[i][routes[i][j]] = true
    		stops[routes[i][j]] = true
    		if routes[i][j] == source{
    			bfs = append(bfs, i)
    			delete(left_buses, i)
    		}
    	}
    }
    _, ok = stops[target]
    if len(bfs) == 0 || !ok{
    	return -1
    }else if source == target{
    	return 0
    }
    delete(stops, source)
    
    for i, j, l = 0, len(bfs), 1; len(bfs) > 0; i++{
    	if i == j{
    		i = 0
    		j = len(bfs)
    		l += 1
    	}
    	_, ok = buses[bfs[0]][target]
    	if ok{
    		return l
    	}
    	for s, _ = range stops{
    	//for s = 0; s < len(routes[bfs[0]]); s++{
    		_, ok = buses[bfs[0]][s]
    		if ok{
    			for b, _ = range left_buses{
    				_, ok = buses[b][s]
    				if ok{
    					bfs = append(bfs, b)
    					delete(left_buses, b)
    				}
    			}
    			delete(stops,s)
    		}
    	}/*
    	for s = 0; s < len(routes[bfs[0]]); s++{
    		_, ok = stops[routes[bfs[0]][s]]
    		if ok{
    			for b, _ = range left_buses{
    				_, ok = buses[b][routes[bfs[0]][s]]
    				if ok{
    					bfs = append(bfs, b)
    					delete(left_buses, b)
    				}
    			}
    			delete(stops, routes[bfs[0]][s])
    		}
    	}*/
    	bfs = bfs[1:]
    }

    return -1
}
