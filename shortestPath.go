/*leetcode 1293*/

package main

import "fmt"

func main(){
	K := []int{1,1,2,1,1,2}
	grids := [][][]int{
		{{0,0,0},{1,1,0},{0,0,0},{0,1,1},{0,0,0}},
		{{0,1,1},{1,1,1},{1,0,0}},
		{{0,1,1},{1,1,1},{1,0,0}},
		{{0,0,0},{1,1,1},{0,0,0},{0,1,1},{0,0,0}},
		{{0,0,0},{1,1,1},{1,0,0},{0,1,1},{0,0,0}},
		{{0,0,0},{1,1,1},{1,0,0},{0,1,1},{0,0,0}},		
	}
	answers := []int{6,-1,4,6,-1,6}
	var j int
	for i := 0; i < len(K); i++{
		j =  shortestPath(grids[i], K[i])
		if answers[i] != j{
			fmt.Println("error", grids[i], K[i], "returning answer", j)
		}else{
			fmt.Println("correct answer")
		}
	}
}

func shortestPath(grid [][]int, k int) int {
      if len(grid) == 1 && len(grid[0]) == 1{
    	return 0
    }
    var i, n, e, xc, yc, dis int //n - number of cells in a level
	n = 1
    bfs := [][]int{{0,0,k}} //coordinate_x, cordinate_y, max number of obstacles elemination possible 
	
	//last_visited stores the value of k(max number of obstacle elimination possible) when the cell was last visited
    last_visited := make([][]int, len(grid)); visited := make([][]bool, len(grid))
    last_visited[0] = make([]int, len(grid[0])); visited[0] = make([]bool, len(grid[0])) 
    
	visited[0][0] = true; last_visited[0][0] = k
	dirs := [][]int{{0,-1},{0,1},{-1,0},{1,0}}
    
	for ; len(bfs) > 0; i++{
    		if i == n{ //new  bfs level; all cells of previous level traversed
    			i = 0
    			n = len(bfs)
    			dis+= 1
    		}
    		for e = 0; e < 4; e++{
    			xc = bfs[0][0]+dirs[e][0]
    			yc = bfs[0][1]+dirs[e][1]
    			if (xc > -1) && (xc < len(grid)) && (yc > -1) && (yc < len(grid[0])){
    				
					if (xc == len(grid)-1 && yc == len(grid[0])-1){return dis+1}
    				if len(visited[xc]) == 0{ //means cell  is not visited yet
    					visited[xc] = make([]bool, len(grid[0])); last_visited[xc] = make([]int, len(grid[0]))
					}

    				if grid[xc][yc] == 1{
    					/*cannot travel to obstacle(xc,yc) bcoz max limit of obstacle elimination reached
    						or
    					  cannot travel to any visited cell with lesser k (max limit of obstacle elimination) 
    					*/
    					if bfs[0][2] == 0 || (visited[xc][yc] && bfs[0][2] < last_visited[xc][yc]+2){
    						continue
    					}
    					bfs = append(bfs, []int{xc, yc, bfs[0][2]-1})
    					visited[xc][yc] = true; last_visited[xc][yc] = bfs[0][2]-1
    				}else if !(visited[xc][yc] && bfs[0][2] <= last_visited[xc][yc]){ //cell not visited or visited with lesser k value
    					bfs = append(bfs, []int{xc, yc, bfs[0][2]})
    					visited[xc][yc] = true; last_visited[xc][yc] = bfs[0][2]
    				}
    			}
    		}
    		bfs = bfs[1:]
    }
    
    return -1
}