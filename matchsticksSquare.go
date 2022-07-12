package main

import "fmt"

func main() {
	var output bool
	tests := [][]int{
		{1, 1, 2, 2, 2},
		{3, 3, 3, 3, 4},
		{1, 1, 1, 1},
		{1, 2, 2, 1, 4, 2},
		{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
	}
	ans := []bool{true, false, true, false, true}
	for i := 0; i < len(tests); i++ {
		output = makesquare(tests[i])
		if output != ans[i] {
			fmt.Println("error at i =", i, "output", output, "expected", ans[i])
			break
		} else {
			fmt.Println("correct")
		}
	}
}

/*leetcode 473- memorization with backtracking*/

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	var i, v int
	for ; i < len(matchsticks); i++ {
		v += matchsticks[i]
	}
	if v%4 != 0 {
		return false
	}
	m := map[int]bool{}
	curr_visited := 0
	curr_sum := 0
	side_len := v / 4
	return rec(0, &curr_sum, &side_len, &curr_visited, &matchsticks, m)
}

func rec(i int, curr_sum, side_len, curr_visited *int, matchsticks *[]int, m map[int]bool) bool {
	*curr_visited |= 1 << i
	*curr_sum += (*matchsticks)[i]

	_, ans := m[*curr_visited]
	if ans {
		ans = m[*curr_visited]
		*curr_visited ^= 1 << i
		*curr_sum -= (*matchsticks)[i]
		return ans
	}
	if *curr_sum == *side_len && (*curr_visited^((1<<len(*matchsticks))-1) == 0) {
		m[*curr_visited] = true
	} else if *curr_sum > *side_len {
		m[*curr_visited] = false
	} else {
		if *curr_sum == *side_len {
			*curr_sum = 0
		}
		m[*curr_visited] = false
		for j := 0; j < len(*matchsticks); j++ {
			if (*curr_visited & (1 << j)) == 0 { // not visited
				if rec(j, curr_sum, side_len, curr_visited, matchsticks, m) {
					m[*curr_visited] = true
					break
				}
			}
		}
	}
	ans = m[*curr_visited]
	*curr_visited ^= 1 << i
	if *curr_sum == 0 {
		*curr_sum = *side_len - (*matchsticks)[i]
	} else {
		*curr_sum -= (*matchsticks)[i]
	}
	return ans
}
