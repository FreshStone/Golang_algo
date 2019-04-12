package main

import "fmt"

func main(){
	a := "000" //232, 8    105, 5    2431 12     23 6      00 0     000 0     1234 0     10002 2
	target := 0
	fmt.Println(addOperators(a, target))
}

func addOperators(num string, target int) []string{
	if len(num) == 0{
		return []string{}
	}
	var res []string
	addops(num, target, 0, -1, +1, "", &res)
	return res
}

func addops(num string, target, prev, mul, sign int, s string, res *[]string){
	if len(num) == 0{
		if mul == -1{
			if target == sign*prev{
				*res = append(*res, s)
			}
		}else{
			if target == sign*mul*prev{
				*res = append(*res, s)
			}
		}
		return
	}

	s = s + string(num[0])
	prev = prev*10 + int(num[0]-48)
	if (len(s) == 1 && num[0] != 48) || (len(s) > 1 && (prev - (prev/10) != 0 || len(num) == 1)){
		addops(num[1:], target, prev, mul, sign, s, res)
	}

	if mul != -1{
		prev = mul*prev
	}

	if len(num) > 1{
		addops(num[1:], target-sign*prev, 0, -1, +1,  s + "+", res)
		addops(num[1:], target-sign*prev, 0, -1, -1,  s + "-", res)
		addops(num[1:], target, 0, prev, sign, s + "*", res)
	}
}
