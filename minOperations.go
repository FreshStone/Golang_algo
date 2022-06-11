package main

import "fmt"

func main() {
	tests := [][]int{
		{2, 3, 1, 1, 1},
		{1, 1, 4, 2, 3},
		{5, 6, 7, 8, 9},
		{3, 2, 20, 1, 1, 3},
		{3, 2, 5, 6, 3, 7},
		{2},
		{4},
	}
	x := []int{5, 5, 4, 10, 26, 3, 3}
	ans := []int{2, 2, -1, 5, 6, -1, -1}
	for i := 0; i < len(tests); i++ {
		if ans[i] != minOperations(tests[i], x[i]) {
			fmt.Println("error") //, "expected", ans[i], "returning", minOperations(tests[i], x[i]))
		} else {
			fmt.Println("corrrect")
		}
	}
}

/* leetcode 1658
2 pointers approach can be used to avoid using map
approach-
	find the maximum length of the contiguous subarray, sum of whose elements is s-x //s = sum of all elements of nums
	return len(nums) -len(contiguous_arr)
*/
func minOperations(nums []int, x int) int { //prefix sum approach using map
	if len(nums) == 1 && nums[0] != x {
		return -1
	}
	var ok bool
	var i, j int
	ans := len(nums) + 1
	sum := make([]int, len(nums)+1)
	m := make(map[int]int)
	for ; i < len(nums); i++ {
		sum[i+1] = nums[i] + sum[i]
	}
	if sum[i] < x {
		return -1
	} else if sum[i] == x {
		return i
	}
	m[0] = 0
	for i = 0; i < len(nums); i++ {
		j, ok = m[x-sum[len(nums)]+sum[i+1]]
		if ok && (ans > j+len(nums)-1-i) {
			ans = j + len(nums) - 1 - i
		}
		m[sum[i+1]] = i + 1
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}

func minOperations_dp(nums []int, x int) int { //using dp
	if len(nums) == 1 && nums[0] != x {
		return -1
	}
	var i, j, curr int
	sum := make([]int, len(nums)+1)
	dp := make([]int, len(sum)+1)
	for ; i < len(nums); i++ {
		sum[i+1] = nums[i] + sum[i]
		dp[i+1] = len(sum)
	}
	dp[i+1] = len(sum)
	if sum[i] < x {
		return -1
	} else if sum[i] == x {
		return i
	}
	for i = 0; i < len(nums); i++ {
		for j = 1; j < len(sum)-i; j++ {
			curr = sum[j-1] + sum[len(nums)] - sum[i+j]
			//fmt.Println(j-1, i+j-1, curr)
			if curr > x {
				dp[j] = len(sum)
			} else if curr == x {
				dp[j] = 0
			} else {
				dp[j] = 1 + min(dp[j], dp[j+1])
			}
		}
		//fmt.Println(dp)
	}
	if dp[1] > len(nums) {
		return -1
	}
	return dp[1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
