package main

import (
	"fmt"
)

	var w = [][]uint{{1, 3, 5}, {2, 5, 6}, {4, 6, 5}, {5, 8, 11}, {6, 7,4}, {7,9,4}}
	//arranged in increasing order of starttime
	//joblist{startime, endtime, value}
        var l = w[0][0] //lowest working time(l+1)

func main(){
	m := func(a [][]uint) uint{
		max := uint(0)
		for i, _ := range a{
			if a[i][1] > max {
				max = a[i][1]
			}
		}
		return max
	}(w)
	dp := make([][]uint, len(w)+1)
	for i, _ := range dp{
		dp[i] = make([]uint, m-l+1)
	}

	for i := 1; i<len(w)+1; i++{
		for j := uint(1); j < m-l+1; j++{
			if w[i-1][1] != l+j{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1], w[i-1][2]+dp[i][w[i-1][0]-l])
			}
		}
	}
	fmt.Println(dp[len(w)][m-l])
	findjobs(dp)
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

func findjobs(dp [][]uint){
	i := len(dp)-1
	j := len(dp[0])-1
	v := dp[i][j]
	for {
		if v == uint(0){
			break
		}
		if (v > dp[i-1][j]) && (v > dp[i][j-1]){
			v = v-w[i-1][2]
			j = int(w[i-1][0]-l)
			fmt.Printf("job index %v selected from job list\n", i)
		}else if v == dp[i-1][j]{
			i = i-1
		}else {
			j = j-1
		}
	}
}
