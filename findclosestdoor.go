/*
"o" - open cell
"w" - wall
"g" - guard
*/
package main

import "fmt"

func main(){
	a := [][]string{{"o","o","o","o","g"},{"o","w","w","o","o"},{"o","o","o","w","o"},{"g","w","w","w","o"},{"o","o","o","o","g"}}
	res := findclosestguard(a)
	for _, v := range res{
		fmt.Println(v)
	}
}

func findclosestguard(a [][]string)[][]int{
	res := make([][]int, len(a))
	for i := 0; i < len(a); i++{
		res[i] = make([]int, len(a[0]))
	}

	type s struct{
		x, y, par int
	}

	var queue []s
	var k, curr int
	row := []int{-1,1,0,0}
	col := []int{0,0,-1,1}

	exist := func(m, n int)bool{
		 if (m > -1 && m < len(a)) && (n > -1 && n < len(a[0])){
			 return true
		 }
		 return false
	 }

	 unvisitedopencell := func(m, n int)bool{
		 if string(a[m][n]) == "o"{
			 if res[m][n] == 0{
				return true
			}
		 }else if string(a[m][n]) == "g"{
			 res[m][n] = 0
		 }else{
			 res[m][n] = -1
		 }
		 return false
	 }


	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[0]); j++{
			if string(a[i][j]) == "o"{
				if res[i][j] == 0{//not visited yet
					res[i][j] = -1
					curr = 0
					queue = append(queue, s{x: i, y: j, par: -1})
					for len(queue) > 0{
						k = 0
						for ;k < 4; k++{
							if exist(queue[curr].x+row[k], queue[curr].y+col[k]){
								if unvisitedopencell(queue[curr].x+row[k], queue[curr].y+col[k]){
									queue = append(queue, s{x: queue[curr].x+row[k], y: queue[curr].y+col[k], par: curr})
									res[queue[curr].x+row[k]][queue[curr].y+col[k]] = -1
									curr = len(queue)-1
									break
								}else if res[queue[curr].x+row[k]][queue[curr].y+col[k]] > -1{
									if res[queue[curr].x][queue[curr].y] == -1{
										res[queue[curr].x][queue[curr].y] = 1+res[queue[curr].x+row[k]][queue[curr].y+col[k]]
									}else{
										res[queue[curr].x][queue[curr].y] = min(res[queue[curr].x][queue[curr].y], 1+res[queue[curr].x+row[k]][queue[curr].y+col[k]])
									}
								}
							}
						}

						if k == 4{//all neighbours visited
							if curr == 0{
								queue = queue[1:]
							}else{
								curr = queue[curr].par
							}
						}
					}
				}
			}else if string(a[i][j]) == "g"{
				res[i][j] = 0
			}else{
				res[i][j] = -1
			}
		}
	}
	return res
}

func min(i, j int)int{
	if i < j{
		return i
	}
	return j
}
