package main

import "fmt"

func main() {
	tests := [][]int{
		{2, 2, 3, 4, 4, 5, 6, 7, 7},
		{1, 2, 3, 3, 4, 5},
		{1, 2, 3, 3, 4, 4, 5, 5},
		{1, 2, 3, 4, 4, 5},
		{1, 2, 3, 4, 4, 5, 5},
		{1, 2, 3, 4, 4, 5, 5, 6},
		{1, 2, 3, 4, 4, 5, 5, 6, 7},
		{1, 2, 3, 5, 5, 6, 7},
		{1, 2, 3, 5, 5, 6, 6},
		{1, 2, 3, 5, 5, 6, 6, 7, 8},
	}
	ans := []bool{false, true, true, false, false, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		if ans[i] != isPossible(tests[i]) {
			fmt.Println("error") //, "expected", ans[i], "returning", minOperations(tests[i], x[i]))
		} else {
			fmt.Println("corrrect")
		}
	}
}
/*leetcode 659; tc- O(n) */
func isPossible(nums []int) bool { //nums is a non decreasing array
	if len(nums) < 3 {
		return false
	}
	var i, j int
	var next []int
	stack := make([]int, 3)
	prev := nums[0]
	nums[0] = 1
	for i = 1; i < len(nums); i++ { //turning nums to freq matrix
		if nums[i] == prev {
			nums[j] += 1
			continue
		} else if nums[i] > prev+1 {
			next = append(next, j+1)
		}
		j += 1
		prev = nums[i]
		nums[j] = 1
	}
	next = append(next, j+1)
//	fmt.Println(next, nums)
	for i, j = 0, 0; j < len(next); j++ {
		if next[j]-i < 3 {
			return false
		}
		for ; i < next[j]; i++ {
			if nums[i] >= stack[2] {
				stack[0], stack[1], stack[2] = stack[1], stack[2], nums[i]
			} else if nums[i] >= stack[2]-stack[0] {
				stack[0], stack[1], stack[2] = stack[1]-stack[2]+nums[i], nums[i], nums[i]
			} else {
				return false
			}
		}
		if !(stack[0] == stack[1] && stack[1] == stack[2]) {
			return false
		}
		stack[0], stack[1], stack[2] = 0, 0, 0
	}
	return true
}
