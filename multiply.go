package main

import (
	"fmt"
	"math"
)

func main(){
	a := "1331789514112297869753139512633768066017431670578465899940305692590313366445604440249153185776128379607342024"
	b := "905175220518204524894521085019700317051204034"
	fmt.Println(multiply(a, b))
}

func multiply(num1 string, num2 string)string{
	var s string
	var a, b, r, cur, carry, leftover, previous int
	var num2_arr, res_arr []int
	for i := 0; len(num1) > 0; i++{
		a, num1 = help(num1)
		for j := 0; j < len(num2_arr) || len(num2) > 0 || previous != 0 || carry != 0; j++{
			if i == 0{
				if len(num2) == 0{
					r = previous
				}else{
					b, num2 = help(num2)
					num2_arr = append(num2_arr, b)
					r = a*b + previous
				}
				leftover = r/(int(math.Pow10(9)))
				res_arr = append(res_arr, r - leftover*int(math.Pow10(9)))
				previous = leftover
			}else{
				if j > len(num2_arr) -1{
					r = previous
				}else{
					b = num2_arr[j]
					r = a*b + previous
				}
                                leftover = r/(int(math.Pow10(9)))
				if cur > len(res_arr)-1{
					res_arr = append(res_arr, r - leftover*int(math.Pow10(9)) + carry)
				}else{
					res_arr[cur] = res_arr[cur] + r - leftover*int(math.Pow10(9)) + carry
				}

				if res_arr[cur] > int(math.Pow10(9))-1{
					carry = res_arr[cur]/int(math.Pow10(9))
					res_arr[cur] = res_arr[cur] - carry*int(math.Pow10(9))
				}else{
					carry = 0
				}
				previous = leftover
				cur += 1
			}
		}
		cur = i+1
	}
	for i, v := range res_arr{
		var no_zero int
		s = tostring(v) + s
		if v <  int(math.Pow10(8)) && i != len(res_arr)-1{
			no_zero = 9 - int(math.Floor(math.Log10(float64(v)))) - 1
			for no_zero > 0{
				s = "0" + s
				no_zero -= 1
			}
		}
	}
	return s
}

func help(num string)(int, string){
		var a int
	        if len(num) < 10{
                        a = tonum(num)
                        num = ""
                }else{
			a = tonum(num[(len(num)-9):])
                        num = num[:(len(num)-9)]
                }
		return a, num
}


func tonum(s string) int{
	var n int
	l := len(s)
	for i := 0; i < l; i++{
		n += int(math.Pow10(i))*int(byte(s[l-1-i])-48)
	}
	return n
}

func tostring(n int) string{
	var s string
	var tmp int
	for n > 0{
		tmp = n/10
		s = string(uint8(n - tmp*10)+48) + s
		n = tmp
	}
	return s
}
