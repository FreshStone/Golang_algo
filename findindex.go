/*
find starting index of all anagrams of s in a
*/
package main

import "fmt"

func main(){
	a := "abxabaa"
	s := "aab"
	fmt.Println(findstartindex(a, s))
}

func findstartindex(a, s string)[]int{
	if len(a) < len(s){
		return []int{}
	}
	store := make(map[string]int)
	for _, v := range s{
		store[string(v)] += 1
	}
	var res []int
	var tmpStore map[string]int
	var curr_len, start int
	var tmpStore_exist bool
	for i, v := range a{
		if store[string(v)] != 0{
			if tmpStore_exist{
				if tmpStore[string(v)] < store[string(v)]{
					curr_len += 1
					tmpStore[string(v)] += 1
				}else{
					//shift start
					for{
						tmpStore[string(a[start])] -= 1
						if string(a[start]) == string(v){
							break
						}
						start += 1
					}
					start += 1
					curr_len = i - start + 1
				}

				if curr_len == len(s){
                                        res = append(res, start)
                                }
			}else{
				tmpStore = make(map[string]int)
				tmpStore[string(v)] = 1
				start = i
				curr_len = 1
				tmpStore_exist = true
			}
		}else{
			tmpStore_exist = false
		}
	}
	return res
}
