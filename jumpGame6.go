package main

import "fmt"

func main() {
	var output int
	tests := [][]int{
		{1, -1, -2, 4, -7, 3},
		{10, -5, -2, 4, 0, 3},
		{1, -5, -20, 4, -1, 3, -6, -3},
	}
	k := []int{2, 3, 2}
	ans := []int{7, 17, 0}
	for i := 0; i < len(tests); i++ {
		output = maxResult(tests[i], k[i])
		if output != ans[i] {
			fmt.Println("error at i =", i, "output", output, "expected", ans[i])
			break
		} else {
			fmt.Println("correct")
		}
	}
}

/*leetcode 1696; tc- O(n); using max dequeue
a queue is maintained to store max elements in sorted order from nums[i-k] to
num[i-1] with max_queue[0] containing the index and value of max element.
for each nums[i] -> elements from back of the queue are popped until a greater val is
found and then nums[i](+max_queue[0][1]) is pushed at the last of the queue
*/
func maxResult(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var i, j, v int
	max_queue := [][]int{{0, nums[0]}}
	for i = 1; i < len(nums); i++ {
		if max_queue[0][0] < i-k { //max_val out of range
			max_queue = max_queue[1:]
		}
		v = nums[i] + max_queue[0][1]
		for j = len(max_queue) - 1; j > -1; j-- {
			if max_queue[j][1] > v {
				break
			}
		}
		max_queue = max_queue[:j+1]
		max_queue = append(max_queue, []int{i, v})
		//fmt.Println(max_queue)
	}
	return max_queue[j+1][1]
}
