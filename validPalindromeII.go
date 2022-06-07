package main

import "fmt"

func main() {
	tests := []string{"aba", "abca", "abc", "cabc", "dbaaaacbd", "cbaaaacbd", "ccaaaaabcc", "eedede"}
	//tests := []string{"dbaaaacbd"}
	ans := []bool{true, true, false, true, true, false, true, true}
	for i := 0; i < len(tests); i++ {
		if ans[i] != validPalindrome(tests[i]) {
			fmt.Println("error", i)
		} else {
			fmt.Println("Correct")
		}
	}
}

func validPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	is_pal := []bool{true, true, true}
	deleted := []bool{false, false, false}
	var i int
	i = (len(s) - 1) / 2
	start := []int{i, i - 1, i}
	for i = (len(s) + 1) / 2; i > 0; i-- {
		if !is_pal[0] && !is_pal[1] && !is_pal[2] {
			return false
		}
		if is_pal[0] {
			if s[start[0]] != s[len(s)-1-start[0]] {
				is_pal[0] = false
			}
			start[0] -= 1
		}
		if is_pal[1] && start[1] >= 0 {
			if deleted[1] {
				if s[start[1]] != s[len(s)-1-start[1]] {
					is_pal[1] = false
				}
				start[1] -= 1
			} else if s[start[1]] != s[len(s)-2-start[1]] {
				deleted[1] = true
			} else {
				start[1] -= 1
			}
		}

		if is_pal[2] {
			if deleted[2] {
				if s[start[2]] != s[len(s)-1-start[2]] {
					is_pal[2] = false
				}
			} else if start[2] > 0 && s[start[2]] != s[len(s)-start[2]] {
				deleted[2] = true
			}
			start[2] -= 1
		}
	}
	if is_pal[0] || is_pal[1] || is_pal[2] {
		return true
	}
	return false
}
