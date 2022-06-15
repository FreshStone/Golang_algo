package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]string{
		{"a", "b", "ba", "bca", "bda", "bdca"},
		{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
		{"abcd", "dbqca"},
	}
	ans := []int{4, 5, 1}
	for i := 0; i < len(tests); i++ {
		if ans[i] != longestStrChain(tests[i]) {
			fmt.Println("error") //, "expected", ans[i], "returning", minOperations(tests[i], x[i]))
		} else {
			fmt.Println("corrrect")
		}
	}
}

func longestStrChain(words []string) int {
	if len(words) == 1 {
		return 1
	}
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	//fmt.Println(words)
	type m map[string]int
	var i, j, v, t, ans int
	var ok bool
	maps := make([]m, len(words[len(words)-1])+1)
	maps[len(words[0])] = map[string]int{}
	for ; i < len(words); i++ {
		if len(words[i]) == len(words[0]) {
			maps[len(words[i])][words[i]] = 1
			ans = 1
		} else {
			t = 1
			if len(words[i]) > len(words[i-1]) {
				maps[len(words[i])] = map[string]int{}
			}
			for j = 0; j < len(words[i]); j++ {
				v, ok = maps[len(words[i])-1][words[i][:j]+words[i][j+1:]]
				if ok {
					t = max(t, 1+v)
				}
			}
			maps[len(words[i])][words[i]] = t
			ans = max(ans, t)
		}
	}
	//fmt.Println(maps)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
