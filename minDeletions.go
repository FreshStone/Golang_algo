package main

import (
	"fmt"
	"sort"
)

func main() {
	var output int
	tests := []string{
		"aabbcccdddddeeeeeffffffgggggghhhhhhhhhiiiiiiiii",
		"aaabbbcccddd",
		"aaabbbcccdddeee",
		"aab",
		"aaabbbcc",
		"ceabaacb",
	}
	ans := []int{9, 6, 9, 0, 2, 2}
	for i := 0; i < len(tests); i++ {
		output = minDeletions(tests[i])
		if ans[i] != output {
			fmt.Println("error Output", output, "expected", ans[i])
		} else {
			fmt.Println("correct")
		}
	}

}

/*leetcode- 1647
greedy approach is the optimal approach
sort(feq)
start iterating from the end
	if feq[j] <= feq[j-1]{
		ans += feq[j-1] - max(0, feq[j]-1)
		feq[j-1] = feq[j]-1
	}
*/
func minDeletions(s string) int {
	if len(s) == 1 {
		return 0
	}
	var i, j, ans int
	var stack [][]int
	feq := make([]int, 26)
	for ; i < len(s); i++ {
		feq[s[i]-97] += 1
	}
	sort.Ints(feq)
	for i = 24; i >= 0; i-- {
		if feq[i] == 0 {
			break
		} else if feq[i+1]-feq[i] > 1 {
			stack = append(stack, []int{feq[i], feq[i+1] - feq[i] - 1})
		}
	}
	i += 1
	stack = append(stack, []int{-100000, feq[i] + 99999})
	for i = 24; i >= 0 && feq[i] > 0; {
		if feq[i] == feq[i+1] {
			if stack[j][0] >= feq[i] || stack[j][1] == 0 {
				j += 1
			} else {
				//	fmt.Println(feq[i], "--> ", max(0, stack[j][0]+stack[j][1]))
				ans += feq[i] - max(0, stack[j][0]+stack[j][1])
				stack[j][1] -= 1
				i -= 1
			}
		} else {
			i -= 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
