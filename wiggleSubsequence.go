package main

import "fmt"

func main() {
	var output int
	tests := [][]int{
		{1, 7, 4, 9, 2, 5},
		{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 4, 7, 2, 5},
		{1, 7, 4, 5, 5},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}
	ans := []int{6, 7, 2, 4, 4, 2, 2}
	for i := 0; i < len(tests); i++ {
		output = wiggleMaxLength(tests[i])
		if output != ans[i] {
			fmt.Println("error at i =", i, "output", output, "expected", ans[i])
			break
		} else {
			fmt.Println("correct")
		}
	}

}
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var inc bool
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] > nums[i-1] {
			if !inc {
				ans += 1
				inc = true
			}
		} else if inc {
			ans += 1
			inc = false
		} else if ans == 1 {
			ans += 1
		}
	}
	return ans
}
