package main

import (
	"fmt"
	"sort"
)

func main() {
	var output []int
	tests_intervals := [][][]int{
		{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
		{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
	}
	tests_queries := [][]int{
		{2, 3, 4, 5},
		{2, 19, 5, 22},
	}
	ans := [][]int{
		{3, 3, 1, 4},
		{2, -1, 4, 6},
	}
	for i := 0; i < len(tests_intervals); i++ {
		output = minInterval(tests_intervals[i], tests_queries[i])
		if checkAnswer(output, ans[i]) {
			fmt.Println("correct answer")
		} else {
			fmt.Println("wrong answer for index", i)
			fmt.Println("output")
			fmt.Println(output)
			fmt.Println("expected")
			fmt.Println(ans[i])
			break
		}
	}
}

/* leetcode 1851- priority queue approach.
1.create queries_util queries_util[i] = []int{id, val]} and sort it by queries_util[i][1]
2. iterate through intervals and for each intervals[i][0] use binary search in queries_util
   to find the lowest index(j) whose value is equal to or greater than intervals[i][0] and append i to
   start[j]
3. for each index i in queries_util
	a. push each value in start[i] to pq(min heap)
	b. pop interval indexes from pq until we find an interval whose end val >= queries[i][1]
	c. update queries[queries_util[i][0]] = length of interval(pq[0])
*/

func checkAnswer(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func minInterval(intervals [][]int, queries []int) []int {
	start := make([][]int, len(queries)+1)
	queries_util := make([][]int, len(queries))
	var i, j, m, l, h int
	var pq []int
	for ; i < len(queries); i++ {
		queries_util[i] = []int{i, queries[i]}
	}
	sort.Slice(queries_util, func(i, j int) bool {
		return queries_util[i][1] < queries_util[j][1]
	})

	length := func(id int) int {
		return intervals[id][1] - intervals[id][0]
	}
	pop := func() {
		if len(pq) == 1 {
			pq = pq[1:]
			return
		}
		l = len(pq) - 1
		pq[0] = pq[l]
		pq = pq[:l]
		for h = 0; 2*h+2 < l; {
			if (length(pq[2*h+1]) <= length(pq[2*h+2])) && (length(pq[2*h+1]) < length(pq[h])) {
				pq[h], pq[2*h+1] = pq[2*h+1], pq[h]
				h = 2*h + 1
			} else if (length(pq[2*h+2]) < length(pq[2*h+1])) && (length(pq[2*h+2]) < length(pq[h])) {
				pq[h], pq[2*h+2] = pq[2*h+2], pq[h]
				h = 2*h + 2
			} else {
				break
			}
		}
		if 2*h+1 < l && (length(pq[2*h+1]) < length(pq[h])) {
			pq[h], pq[2*h+1] = pq[2*h+1], pq[h]
		}
		return
	}
	push := func(id int) {
		l = len(pq)
		pq = append(pq, id)
		for l > 0 {
			if length(pq[l]) < length(pq[(l-1)/2]) {
				pq[l], pq[(l-1)/2] = pq[(l-1)/2], pq[l]
				l = (l - 1) / 2
				continue
			}
			break
		}
		return
	}
	for i = 0; i < len(intervals); i++ {
		j = intervals[i][0]
		l, h = 0, len(queries)-1
		for l <= h {
			m = (l + h) / 2
			if j > queries_util[m][1] {
				l = m + 1
			} else {
				h = m - 1
			}
		}
		start[l] = append(start[l], i)
	}
	for i = 0; i < len(queries); i++ {
		for j = 0; j < len(start[i]); j++ {
			push(start[i][j])
		}
		for len(pq) > 0 {
			if intervals[pq[0]][1] >= queries_util[i][1] {
				break
			}
			pop()
		}
		//fmt.Println(queries_util[i], pq)
		if len(pq) == 0 {
			queries[queries_util[i][0]] = -1
		} else {
			queries[queries_util[i][0]] = intervals[pq[0]][1] - intervals[pq[0]][0] + 1
		}

	}
	return queries
}
