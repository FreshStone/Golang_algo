package main

import "fmt"

func main() {
	tests := []int{4, 5, 6, 7}
	ans := []int{2, 10, 4, 40}
	for i := 0; i < len(tests); i++ {
		if ans[i] != totalNQueens(tests[i]) {
			fmt.Println("error", i, totalNQueens(tests[i]))
		} else {
			fmt.Println("Correct Answer")
		}
	}
}

func totalNQueens(n int) int {
	ans := []int{0, 1, 0, 0}
	if n < 4 {
		return ans[n]
	}
	var i int
	positions := make([]int, n) //postions[i] = position of queen in the ith row
	for i = 0; i < (n / 2); i++ {
		positions[0] = i
		ans[0] += rec(positions, 1) //postions, row_number
	}
	if n%2 == 1 {
		positions[0] = i
		return 2*ans[0] + rec(positions, 1)
	}
	return 2 * ans[0]
}

func rec(pos []int, r int) int {
	if r == len(pos) {
		return 1
	}
	var i, ans int
	m := map[int]bool{}
	for ; i < len(pos); i++ {
		m[i] = false
	}

	for i = 0; i < r; i++ {
		delete(m, pos[i])
		delete(m, pos[i]-r+i)
		delete(m, pos[i]+r-i)
	}

	for i, _ = range m {
		pos[r] = i
		ans += rec(pos, r+1)
	}
	return ans
}
