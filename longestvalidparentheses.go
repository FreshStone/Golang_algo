package main

import "fmt"

func main(){
//	s := "(()())"
	s := ")()())()()("
	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	if len(s) == 0{
		return 0
	}
	var max , current int
	type node struct{
		parent, current_max, last_valid int
	}
	info := make([]node, len(s))
	previous, valid := -1, -1
	for i, v := range s{
		if v == 40{
			info[i] = node{parent: previous, last_valid: valid}
			previous = i
		}else{
			if previous == -1{
				valid = -1
				info[i] = node{parent: previous, current_max: 0, last_valid: valid}
			}else{
				if valid >0 && info[previous].parent == info[info[valid].parent].parent{
					current = info[valid].current_max + 2
				}else if valid > 0 && previous == info[info[valid].parent].parent{
					if info[previous].last_valid > 0 && previous - info[previous].last_valid == 1{
						current = info[info[previous].last_valid].current_max +info[valid].current_max+2
					}else{
						current = info[valid].current_max + 2
					}
				}else{
					current = 2
				}
				info[i] = node{previous, current, valid}
				previous = info[previous].parent
				valid = i
				if current > max{
					max = current
				}
			}
		}
	}
	return max
}

