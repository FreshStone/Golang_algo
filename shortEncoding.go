package main

import "fmt"

func main() {
	tests := [][]string{
		{"time"},
		{"time", "me", "bell"},
		{"time", "me", "bell", "abime", "anime"},
		{"time", "me", "abtime", "anime", "bell"},
	}
	ans := []int{5, 10, 22, 18}
	for i := 0; i < len(tests); i++ {
		if ans[i] != minimumLengthEncoding(tests[i]) {
			fmt.Println("error output", minimumLengthEncoding(tests[i]), "expected", ans[i])
		} else {
			fmt.Println("correct")
		}
	}
}

/* leetcode 820 -
also see lee215's solution
*/
func minimumLengthEncoding(words []string) int {
	if len(words) == 1 {
		return len(words[0]) + 1
	}
	type node struct {
		m   map[byte]*node
		end bool
	}
	var tmp, n *node
	var ok, new_branch bool
	var i, j, l, ans int
	root := &node{m: map[byte]*node{}}
	for ; i < len(words); i++ {
		l = 0
		new_branch = false
		n = root
		for j = len(words[i]) - 1; j >= 0; j-- {
			tmp, ok = n.m[words[i][j]]
			if !ok {
				if !new_branch {
					if n.end {
						l = 0
						n.end = false
					} else {
						l += 1
					}
				}
				new_branch = true
				tmp = &node{m: map[byte]*node{}}
				n.m[words[i][j]] = tmp
			}
			l += 1
			n = tmp
		}
		if new_branch {
			n.end = true
			ans += l
		}
	}
	return ans
}
