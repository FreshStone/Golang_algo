package main

import "fmt"

func main() {
	var output int
	tests := [][]int{
		{1, 0, 2},
		{1, 2, 2},
		{1, 2, 2, 2},
		{4, 3, 3, 3, 2, 2, 2, 1, 1, 1, 1},
		{1, 6, 10, 8, 7, 3, 2},
	}
	ans := []int{5, 4, 5, 14, 18}
	for i := 0; i < len(tests); i++ {
		output = candy(tests[i])
		if output != ans[i] {
			fmt.Println("error at i =", i, "output", output, "expected", ans[i])
			break
		} else {
			fmt.Println("correct")
		}
	}

}

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	stack := []int{1, 0}
	top := 0
	ans := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && top == 1 {
			if stack[0] <= stack[1] {
				ans += stack[1] - stack[0] + 1
			}
			top = 0
			stack[0] = 1
		} else if ratings[i] < ratings[i-1] && top == 0 {
			top = 1
			stack[top] = 0
		} else if ratings[i] == ratings[i-1] {
			if top == 1 && stack[0] <= stack[1] {
				ans += stack[1] - stack[0] + 1
			}
			top = 0
			stack[top] = 0
		}
		stack[top] += 1
		ans += stack[top]
	}
	if top == 1 && stack[0] <= stack[1] {
		ans += stack[1] - stack[0] + 1
	}
	return ans
}
