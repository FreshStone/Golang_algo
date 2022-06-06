package main

import "fmt"

func main() {
	tests := []int{1, 2, 3, 4, 5, 6, 7}
	ans := []int{1, 0, 0, 2, 10, 4, 40}
	for i := 0; i < len(tests); i++ {
		if ans[i] != totalNQueens(tests[i]) {
			fmt.Println("error", i, totalNQueens(tests[i]))
		} else {
			fmt.Println("Correct Answer")
		}
	}
}

func totalNQueens(N int) int {
	ret := []int{0, 1, 0, 0}
	if N < 4 {
		return ret[N]
	}
	var i, j, row, col, r, ans, v, a int
	positions := make([][]int, N)
	last_visited := make([]int, N)
	stack := []int{-1}
	for ; i < N; i++ {
		positions[i] = make([]int, N)
		last_visited[i] = -1
	}

	mark := func() {
		for i = row + 1; i < N; i++ {
			positions[i][col] += a
			if col-i+row >= 0 {
				positions[i][col-i+row] += a
			}
			if col+i-row < N {
				positions[i][col+i-row] += a
			}
		}
	}
	for j = 0; j < (N/2)+1; j++ {
		if (j == N/2) && N%2 == 0 {
			break
		}
		stack = append(stack, 0)
		row, col, a = 0, j, 1
		mark()
		for len(stack) > 1 {
			r = len(stack) - 1
			if r == N {
				v += 1
				stack = stack[:r]
				continue
			}
			if last_visited[r] != -1 {
				row, col, a = r, last_visited[r], -1
				mark()
			}
			last_visited[r] = -1
			if stack[r] == N {
				stack = stack[:r]
				continue
			}
			if positions[r][stack[r]] == 0 { //safe position
				row, col, a = r, stack[r], 1
				mark()
				last_visited[r] = stack[r]
				stack[r] += 1
				stack = append(stack, 0)
			} else {
				stack[r] += 1
			}
		}
		if j == N/2 {
			ans += v
			break
		}
		ans += 2 * v
		v = 0
		row, col, a = 0, j, -1
		mark()
	}
	return ans
}

func totalNQueens_rec(n int) int {
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
