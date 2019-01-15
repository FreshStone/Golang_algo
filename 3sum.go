package main

import (
	"fmt"
	"sort"
)

func main(){
//	a := []int{-1,0,1,2,-1,-4}
//	a := []int{-2,0,0,2,2}
	a := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}
	fmt.Println(threesum(a))
}

func threesum(a []int) [][]int{
	sort.Ints(a)
	var res [][]int
	var j, k int
	for i, v := range a{
		if i > 0 && v == a[i-1]{ //skip evaluating sum for duplicate values
			continue
		}
		j = i+1; k = len(a)-1
		for j < k{
                        if k == i{
                                k -= 1
                                continue
			}else if (j != i+1 && a[j] == a[j-1]) && (k < len(a)-1 && a[k] == a[k+1]){//skip duplicate values while calculating sum
				j += 1
				k -= 1
				continue
			}

                        if a[j] + a[k] == -v{
                                res = append(res, []int{v, a[j], a[k]})
                                j +=1
                                k -=1
                        }else if a[j] + a[k] < -v{
                                j += 1
                        }else{
                                k -= 1
                        }
                }
        }
        return res
}
