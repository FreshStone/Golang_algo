package main

import "fmt"

func main(){
	tests := []int{4}
	k := []int{2,2}
	ans := []int{6,10}
	for i := 0 ; i < len(tests); i++{
		if ans[i] != paintFence(tests[i], k[i]){
			fmt.Println("error", "expected-", ans[i], "returning-", paintFence(tests[i],k[i]))
		}else{
			fmt.Println("correct answer")
		}
	}
}

func paintFence(n, k int)int{ //n >= 1; k >= 1
	if n == 1{
		return k
	}else if k == 1{
		return 0
	}
	
	l_prime := 1000000000+7
	var i int
	dp := []int{k, k*k}
	for i = 2; i < n; i++{
		dp[1], dp[0] = (dp[1]+dp[0])*(k-1)%l_prime, dp[1]
		//fmt.Println(dp[0], dp[1])
	}
	return dp[1]
}