package main

import (
        "fmt"
)

        var w = [][]uint{{1, 3, 5}, {2, 5, 6}, {4, 6, 5}, {6, 7, 4}, {5, 8,11}, {7,9,2}}
	//arranged in incrasing order of endtime

func main(){
	dp := make([]uint, len(w)+1)
	for i := 1; i < len(w)+1; i++{
		if i == 1 {
			dp[i] = w[i-1][2]
		}else {
			dp[i] = max(dp[i-1], (dp[free_slot(i-1, w)]+w[i-1][2]))
		}
	}
	fmt.Println(dp[len(w)])
//	find_slot(dp)
}

func max(nums...uint) uint{
        max := uint(0)
        for i, _ := range nums {
                if nums[i] > max {
                        max = nums[i]
                }
        }
        return max
}


func free_slot(i int, w [][]uint) int {
	a := w[i][0]
	l := -1
	Loop:
		for j := i-1; j >= 0; j-- {
			if a >= w[j][1] {
				l = j
				break Loop
			}
		}
	return l+1
}
/*
find_slot(dp [][]uint) {
}
*/
