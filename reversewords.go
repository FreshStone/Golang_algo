/* reverse words in the string while maintaining relative order of delimiters*/
package main

import "fmt"

func main(){
	a := "//how-you--doing:broda/"
	fmt.Println(reversewords(a))
}

func reversewords(a string)string{
	var ls int
	var ld_started, rd_started, l_change, r_change bool
	re := len(a)-1
	if !isalphabet(a[0]){
		ld_started = true
	}
	if !isalphabet(a[re]){
		rd_started = true
	}

	l := 1
	r := re-1
	for l <= r{
		if isalphabet(a[l]){
			if ld_started{
				ls = l
				ld_started = false
			}
			l += 1
		}else{
			if ld_started{
				l += 1
				l_change = false
			}else{
				l_change = true
			}
		}

		if isalphabet(a[r]){
                        if rd_started{
                                re = r
                                rd_started = false
                        }
			r -= 1
                }else{
                        if rd_started{
                                r -= 1
                                r_change = false
                        }else{
                                r_change = true
                        }
                }

		if l_change && r_change{
			a = string(a[:ls]) + string(a[r+1:re+1]) + string(a[l:r+1]) + string(a[ls:l]) + string(a[re+1:])
			l, r = ls + re - r, re - l + ls
			ld_started = true
			rd_started = true
		}
	}
	return a
}

func isalphabet(v byte)bool{
	if (v > 64 && v < 91) || (v > 96 && v < 123){
		return true
	}
	return false
}
