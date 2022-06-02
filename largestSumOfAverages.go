package main

import "fmt"

func main(){
	nums := [][]int{
		{9,1,2,3,9},
		{9,1,2,3,9},
	}
	k := []int{3,2}
	ans := []float64{20,12.75}
	for i := 0 ; i < len(nums); i++{
		if ans[i] != largestSumOfAverages(nums[i], k[i]){
			fmt.Println("error", "expected-", ans[i], "returning-", largestSumOfAverages(nums[i],k[i]))
		}else{
			fmt.Println("correct answer")
		}
	}
}

func largestSumOfAverages(nums []int, k int) float64 {
    if len(nums) == 1{
        return float64(nums[0])
    }
    var i, low, high, j int
    var ans float64
    sum := make([]float64, len(nums))
    sum[0] = float64(nums[0])
 	for i =1 ; i < len(nums); i++{
 		sum[i] = sum[i-1] + float64(nums[i])
 	}
 	if k == 1{
 		return sum[len(sum)-1]/float64(len(nums))
 	}
 	ans = sum[len(sum)-1]/float64(len(nums))
 	// dp[i][j] = largetst sum of avg's of sub array a[i:] with j+1 max partitions
    dp := make([][]float64, len(nums))    
    for i = len(nums)-1; i > 0; i--{
    	high = min(k-1, len(nums)-i) //maximum possible partitions for sub array a[i:]
    	dp[i] = make([]float64, high)
    	dp[i][0] = (sum[len(sum)-1]-sum[i-1])/float64(len(sum)-i) 
    	for low = 1; low < high; low++{ //low = no of cuts; low+1 = no of partitions
    		for j = i+1; j < len(nums) && low <= len(dp[j]); j++{
    			dp[i][low] = max(dp[i][low], (sum[j-1]-sum[i-1])/float64(j-i) + dp[j][low-1])
    		}
    		dp[i][low] = max(dp[i][low-1], dp[i][low])
    	}
    	//fmt.Println(dp[i])
    }

    for i = 1; i < len(nums); i++{
    	ans = max(ans, sum[i-1]/float64(i) + dp[i][len(dp[i])-1])
    }
    return ans
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}

func max(a, b float64)float64{
	if a > b{
		return a
	}
	return b
}